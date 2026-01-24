package services

import (
	"context"
	"fmt"
	"jm-CICO/config"
	"jm-CICO/models"
	"jm-CICO/utils"
	"log"
	"strings"
	"sync"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// Cache Structure
type CacheEntry struct {
	Data   []models.Entry
	Expiry time.Time
}

var (
	cache      = make(map[string]CacheEntry)
	cacheMutex sync.RWMutex
	cacheTTL   = 5 * time.Minute
)

// Helper to convert row of interface{} to []string
func convertToStrings(row []interface{}) []string {
	res := make([]string, len(row))
	for i, cell := range row {
		res[i] = fmt.Sprintf("%v", cell)
	}
	return res
}

func FetchData(config config.LocationConfig) ([]models.Entry, error) {
	// 1. CACHE CHECK
	cacheKey := config.DataRange // Using DataRange as unique key for the cache

	cacheMutex.RLock()
	if entry, found := cache[cacheKey]; found {
		if time.Now().Before(entry.Expiry) {
			cacheMutex.RUnlock()
			return entry.Data, nil
		}
	}
	cacheMutex.RUnlock()

	// 2. FETCH FROM GOOGLE (Cache Miss or Expired)
	sheetID := "1mBRaSCNk1TTaeJ6AJDhO1kEb7tAJSFZ1ZPhW13nVjq4"

	// Setup Context and Service
	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile("service-account.json"))
	if err != nil {
		return nil, fmt.Errorf("gagal membuat client sheets: %v", err)
	}

	// Define all ranges to fetch in order
	ranges := []string{
		fmt.Sprintf("REPORT!%s", config.DataRange),             // 0
		fmt.Sprintf("REPORT!%s", config.AnggaranRange),         // 1
		fmt.Sprintf("REPORT!%s", config.NopolRange),            // 2
		fmt.Sprintf("REPORT!%s", config.PetugasRange),          // 3
		fmt.Sprintf("REPORT!%s", config.CheckinJmlBulanRange),  // 4
		fmt.Sprintf("REPORT!%s", config.CheckinRupiahRange),    // 5
		fmt.Sprintf("REPORT!%s", config.CheckoutNopolRange),    // 6
		fmt.Sprintf("REPORT!%s", config.CheckoutJmlBulanRange), // 7
		fmt.Sprintf("REPORT!%s", config.CheckoutRupiahRange),   // 8
	}

	// Execute BATCH GET
	resp, err := srv.Spreadsheets.Values.BatchGet(sheetID).Ranges(ranges...).Do()
	if err != nil {
		return nil, fmt.Errorf("gagal fetch batch data: %v", err)
	}

	// Map responses back (Order must match the ranges slice above)
	if len(resp.ValueRanges) < 9 {
		return nil, fmt.Errorf("response google sheets tidak lengkap (kurang dari 9 range)")
	}

	mainDataRaw := resp.ValueRanges[0].Values
	anggaranDataRaw := resp.ValueRanges[1].Values
	nopolDataRaw := resp.ValueRanges[2].Values
	petugasDataRaw := resp.ValueRanges[3].Values
	jmlBlnDataRaw := resp.ValueRanges[4].Values
	rupiahDataRaw := resp.ValueRanges[5].Values
	checkoutNopolDataRaw := resp.ValueRanges[6].Values
	checkoutJmlBlnDataRaw := resp.ValueRanges[7].Values
	checkoutRupiahDataRaw := resp.ValueRanges[8].Values

	var entries []models.Entry

	// Loop through Main Data
	for i, rowRaw := range mainDataRaw {
		row := convertToStrings(rowRaw)

		// Validate bounds against ALL other datasets to prevent panic
		if utils.IsEmptyRow(row) || utils.IsHeaderRow(row) ||
			i >= len(anggaranDataRaw) ||
			i >= len(petugasDataRaw) ||
			i >= len(nopolDataRaw) ||
			i >= len(jmlBlnDataRaw) ||
			i >= len(rupiahDataRaw) ||
			i >= len(checkoutNopolDataRaw) ||
			i >= len(checkoutJmlBlnDataRaw) ||
			i >= len(checkoutRupiahDataRaw) {
			continue
		}

		entry, err := utils.ParseEntry(row)
		if err != nil {
			log.Printf("failed parse row: %v", err)
			continue
		}

		// --- OVERWRITE FIELDS WITH COLUMN DATA ---

		// Petugas (Column O->V, joined)
		petugasRow := convertToStrings(petugasDataRaw[i])
		fullName := ""
		for _, part := range petugasRow {
			p := strings.TrimSpace(part)
			if p != "" {
				fullName += p + " "
			}
		}
		entry.Petugas = strings.TrimSpace(fullName)

		// Anggaran
		anggaranRow := convertToStrings(anggaranDataRaw[i])
		entry.Anggaran = utils.ParseNumber(strings.Join(anggaranRow, ""))

		// Nopol Checkin
		nopolRow := convertToStrings(nopolDataRaw[i])
		entry.Checkin.Nopol = utils.ParseInt(strings.Join(nopolRow, ""))

		// Jml Bulan Checkin
		jmlBlnRow := convertToStrings(jmlBlnDataRaw[i])
		entry.Checkin.JumlahBulan = utils.ParseInt(strings.Join(jmlBlnRow, ""))

		// Rupiah Checkin
		rupiahRow := convertToStrings(rupiahDataRaw[i])
		entry.Checkin.Rupiah = utils.ParseNumber(strings.Join(rupiahRow, ""))

		// Checkout Nopol
		coNopolRow := convertToStrings(checkoutNopolDataRaw[i])
		entry.Checkout.Nopol = utils.ParseInt(strings.Join(coNopolRow, ""))

		// Checkout Jml Bulan
		coJmlBlnRow := convertToStrings(checkoutJmlBlnDataRaw[i])
		entry.Checkout.JumlahBulan = utils.ParseInt(strings.Join(coJmlBlnRow, ""))

		// Checkout Rupiah
		coRupiahRow := convertToStrings(checkoutRupiahDataRaw[i])
		entry.Checkout.Rupiah = utils.ParseNumber(strings.Join(coRupiahRow, ""))

		// Calculate Diff
		entry.SelisihCICO.Nopol = entry.Checkout.Nopol - entry.Checkin.Nopol
		entry.SelisihCICO.JumlahBulan = entry.Checkout.JumlahBulan - entry.Checkin.JumlahBulan
		entry.SelisihCICO.Rupiah = entry.Checkout.Rupiah - entry.Checkin.Rupiah

		entries = append(entries, entry)
	}

	// 3. SAVE TO CACHE
	cacheMutex.Lock()
	cache[cacheKey] = CacheEntry{
		Data:   entries,
		Expiry: time.Now().Add(cacheTTL),
	}
	cacheMutex.Unlock()

	return entries, nil
}

func CalculateSubtotal(entries []models.Entry) models.Subtotal {
	var st models.Subtotal

	for _, entry := range entries {
		st.Anggaran += entry.Anggaran
		st.CheckinNopol += entry.Checkin.Nopol
		st.CheckinJmlBln += entry.Checkin.JumlahBulan
		st.CheckinRupiah += entry.Checkin.Rupiah
		st.CheckoutNopol += entry.Checkout.Nopol
		st.CheckoutJmlBln += entry.Checkout.JumlahBulan
		st.CheckoutRupiah += entry.Checkout.Rupiah
		st.SelisihCICONopol += entry.SelisihCICO.Nopol
		st.SelisihCICOJmlBln += entry.SelisihCICO.JumlahBulan
		st.SelisihCICORupiah += entry.SelisihCICO.Rupiah
	}
	return st
}
