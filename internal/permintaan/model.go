package permintaan

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// StringArray adalah custom type untuk menyimpan []string ke kolom TEXT[] PostgreSQL.
type StringArray []string

func (s StringArray) Value() (driver.Value, error) {
	if s == nil {
		return "{}", nil
	}
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	// Simpan sebagai JSON string; gunakan JSONB di DB, atau konversi ke array literal
	// Jika kolom DB adalah TEXT[], gunakan format: {"val1","val2"}
	result := "{"
	for i, v := range s {
		if i > 0 {
			result += ","
		}
		b, _ := json.Marshal(v)
		result += string(b)
	}
	result += "}"
	_ = b
	return result, nil
}

func (s *StringArray) Scan(src interface{}) error {
	if src == nil {
		*s = StringArray{}
		return nil
	}
	var str string
	switch v := src.(type) {
	case string:
		str = v
	case []byte:
		str = string(v)
	default:
		return fmt.Errorf("StringArray.Scan: unsupported type %T", src)
	}
	// Parse PostgreSQL array literal: {"a","b","c"} atau {}
	if str == "{}" || str == "" {
		*s = StringArray{}
		return nil
	}
	// Trim braces
	str = str[1 : len(str)-1]
	if str == "" {
		*s = StringArray{}
		return nil
	}
	// Simple split by comma — works for simple strings without commas
	var result []string
	if err := json.Unmarshal([]byte("["+str+"]"), &result); err != nil {
		// fallback: split by comma
		parts := splitPGArray(str)
		*s = parts
		return nil
	}
	*s = result
	return nil
}

func splitPGArray(s string) []string {
	// Handles "\"a\",\"b\"" format from PostgreSQL
	var result []string
	var cur string
	inQuote := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '"' {
			inQuote = !inQuote
		} else if c == ',' && !inQuote {
			result = append(result, cur)
			cur = ""
		} else {
			cur += string(c)
		}
	}
	if cur != "" {
		result = append(result, cur)
	}
	return result
}

type Permintaan struct {
	ID                string      `json:"id"`
	PemdaID           string      `json:"pemda_id"`
	AplikasiID        string      `json:"aplikasi_id"`
	Menu              string      `json:"menu"`
	KondisiAwal       string      `json:"kondisi_awal"`
	KondisiDiharapkan string      `json:"kondisi_diharapkan"`
	TanggalPesanan    time.Time   `json:"tanggal_pesanan"`
	TanggalDeadline   time.Time   `json:"tanggal_deadline"`
	Lampiran          StringArray `json:"lampiran"` // URL S3, max 3
	Status            string      `json:"status"`   // proses, selesai, revisi
	CreatedBy         string      `json:"created_by"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

