# Verum Research SDK

> Open-source Go SDK for anonymous peptide research. Auditable anonymization. Verifiable statistics. Shared research pool.

## What This Is

A standalone Go module that any health app can integrate to participate in crowdsourced peptide research. The SDK handles:

- **Compound catalog** — 78+ peptide definitions with dosing, mechanisms, and references
- **Protocol types** — Standardized data structures for peptide protocols and outcomes
- **Biomarker definitions** — 80+ lab biomarkers with LOINC codes and reference ranges
- **Anonymization pipeline** — Auditable, enforced de-identification with k-anonymity
- **Statistical analysis** — Effect sizes, confidence intervals, p-values, dose-response
- **Research API client** — Submit anonymized contributions, query published findings

## Why Open Source

Trust requires auditability. The anonymization pipeline is the single path to contributing data — and anyone can read the code to verify that PII cannot leak. The statistical methods are transparent and reproducible.

## Install

```bash
go get github.com/teslashibe/verum-research-sdk
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/teslashibe/verum-research-sdk/anonymize"
    "github.com/teslashibe/verum-research-sdk/compounds"
    "github.com/teslashibe/verum-research-sdk/contribute"
    "github.com/teslashibe/verum-research-sdk/protocols"
)

func main() {
    compound, _ := compounds.Default().FindByName("BPC-157")

    protocol := protocols.Protocol{
        CompoundID: compound.ID,
        Dose:       250,
        DoseUnit:   "mcg",
        Frequency:  "twice_daily",
        Route:      "subcutaneous",
        StartDate:  time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
        EndDate:    time.Date(2026, 3, 15, 0, 0, 0, 0, time.UTC),
        Goal:       "injury_recovery",
    }

    outcomes := protocols.OutcomeDeltas{
        HRVPctChange:      14.2,
        RecoveryPctChange: 18.5,
        RestingHRChange:   -3.1,
    }

    demographics := anonymize.Demographics{
        AgeRange: "40-49",
        Sex:      "male",
        BMIRange: "23-25",
    }

    contribution, err := anonymize.Prepare(protocol, outcomes, demographics, anonymize.DefaultConfig())
    if err != nil {
        panic(err)
    }

    client := contribute.NewClient("https://api.verum.research", "your-api-key")
    if err := client.Submit(context.Background(), contribution); err != nil {
        panic(err)
    }

    fmt.Println("Contributed anonymously to peptide research")
}
```

## Architecture

```
┌──────────────┐  ┌──────────────┐  ┌──────────────┐
│   Verum App  │  │  Your App    │  │  Other Apps  │
└──────┬───────┘  └──────┬───────┘  └──────┬───────┘
       │                 │                 │
       ▼                 ▼                 ▼
┌─────────────────────────────────────────────────┐
│     verum-research/sdk (this repo)              │
│  compounds · protocols · biomarkers             │
│  anonymize · stats · contribute · report        │
└──────────────────────┬──────────────────────────┘
                       │
                       ▼
         ┌─────────────────────────┐
         │  Verum Research API     │
         │  (shared research pool) │
         └─────────────────────────┘
```

## Packages

| Package | Purpose |
|---|---|
| `compounds` | Peptide compound catalog — 78+ compounds with mechanisms, dosing, references |
| `protocols` | Protocol and outcome data structures, validation, delta computation |
| `biomarkers` | Lab biomarker definitions — 80+ markers with LOINC codes and reference ranges |
| `anonymize` | Anonymization pipeline — bucketing, k-anonymity, differential privacy, leak validation |
| `stats` | Statistical analysis — effect sizes, confidence intervals, p-values, dose-response |
| `contribute` | Research API client — submit contributions, query findings |
| `report` | Report templates — academic structure, disclaimers, dataset export |
| `schema` | Schema versioning for forward compatibility |

## License

Apache 2.0 — see [LICENSE](LICENSE)
