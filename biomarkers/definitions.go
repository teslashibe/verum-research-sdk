package biomarkers

func ptr(v float64) *float64 { return &v }

var Definitions = []Biomarker{
	// ── Iron Panel ──
	{Name: "Iron, Total", LOINCCode: "2498-4", Category: CategoryIron, Unit: "mcg/dL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(50), RefRangeHigh: ptr(180), PeptideRelevance: "oxygen transport, recovery capacity", Aliases: []string{"IRON, TOTAL", "serum iron"}},
	{Name: "Iron Binding Capacity", LOINCCode: "2500-7", Category: CategoryIron, Unit: "mcg/dL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(250), RefRangeHigh: ptr(425), Aliases: []string{"TIBC", "IRON BINDING CAPACITY"}},
	{Name: "Iron Saturation", LOINCCode: "2502-3", Category: CategoryIron, Unit: "%", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(20), RefRangeHigh: ptr(48), Aliases: []string{"% SATURATION", "transferrin saturation"}},
	{Name: "Ferritin", LOINCCode: "2276-4", Category: CategoryIron, Unit: "ng/mL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(38), RefRangeHigh: ptr(380), PeptideRelevance: "iron stores, inflammation marker", Aliases: []string{"FERRITIN"}},

	// ── Lipid Panel ──
	{Name: "Cholesterol, Total", LOINCCode: "2093-3", Category: CategoryLipid, Unit: "mg/dL", RefRangeType: RefRangeLessThan, RefRangeHigh: ptr(200), Aliases: []string{"CHOLESTEROL, TOTAL", "total cholesterol"}},
	{Name: "HDL Cholesterol", LOINCCode: "2085-9", Category: CategoryLipid, Unit: "mg/dL", RefRangeType: RefRangeGreaterThan, RefRangeLow: ptr(40), Aliases: []string{"HDL CHOLESTEROL", "HDL-C", "HDL"}},
	{Name: "Triglycerides", LOINCCode: "2571-8", Category: CategoryLipid, Unit: "mg/dL", RefRangeType: RefRangeLessThan, RefRangeHigh: ptr(150), Aliases: []string{"TRIGLYCERIDES"}},
	{Name: "LDL Cholesterol", LOINCCode: "13457-7", Category: CategoryLipid, Unit: "mg/dL", RefRangeType: RefRangeLessThan, RefRangeHigh: ptr(100), Aliases: []string{"LDL-CHOLESTEROL", "LDL-C", "LDL"}},
	{Name: "Non-HDL Cholesterol", LOINCCode: "43396-1", Category: CategoryLipid, Unit: "mg/dL", RefRangeType: RefRangeLessThan, RefRangeHigh: ptr(130), Aliases: []string{"NON HDL CHOLESTEROL"}},
	{Name: "Cholesterol/HDL Ratio", LOINCCode: "9830-1", Category: CategoryLipid, Unit: "ratio", RefRangeType: RefRangeLessThan, RefRangeHigh: ptr(5.0), Aliases: []string{"CHOL/HDLC RATIO"}},
	{Name: "Lipoprotein(a)", LOINCCode: "10835-7", Category: CategoryCardiac, Unit: "nmol/L", RefRangeType: RefRangeTiered, OptimalHigh: ptr(75), PeptideRelevance: "genetic cardiovascular risk", Aliases: []string{"LIPOPROTEIN (a)", "Lp(a)"}},
	{Name: "Apolipoprotein B", LOINCCode: "1884-6", Category: CategoryCardiac, Unit: "mg/dL", RefRangeType: RefRangeTiered, OptimalHigh: ptr(90), PeptideRelevance: "true atherogenic particle count", Aliases: []string{"APOLIPOPROTEIN B", "ApoB"}},

	// ── Comprehensive Metabolic Panel ──
	{Name: "Glucose", LOINCCode: "2345-7", Category: CategoryMetabolic, Unit: "mg/dL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(65), RefRangeHigh: ptr(99), PeptideRelevance: "GH secretagogues can raise fasting glucose", Aliases: []string{"GLUCOSE", "fasting glucose"}},
	{Name: "BUN", LOINCCode: "3094-0", Category: CategoryKidney, Unit: "mg/dL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(7), RefRangeHigh: ptr(25), Aliases: []string{"UREA NITROGEN (BUN)", "urea nitrogen", "blood urea nitrogen"}},
	{Name: "Creatinine", LOINCCode: "2160-0", Category: CategoryKidney, Unit: "mg/dL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(0.60), RefRangeHigh: ptr(1.29), PeptideRelevance: "kidney safety monitoring during protocols", Aliases: []string{"CREATININE"}},
	{Name: "eGFR", LOINCCode: "33914-3", Category: CategoryKidney, Unit: "mL/min/1.73m2", RefRangeType: RefRangeGreaterThan, RefRangeLow: ptr(60), Aliases: []string{"EGFR", "estimated glomerular filtration rate"}},
	{Name: "Sodium", LOINCCode: "2951-2", Category: CategoryElectrolyte, Unit: "mmol/L", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(135), RefRangeHigh: ptr(146), Aliases: []string{"SODIUM"}},
	{Name: "Potassium", LOINCCode: "2823-3", Category: CategoryElectrolyte, Unit: "mmol/L", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(3.5), RefRangeHigh: ptr(5.3), Aliases: []string{"POTASSIUM"}},
	{Name: "Chloride", LOINCCode: "2075-0", Category: CategoryElectrolyte, Unit: "mmol/L", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(98), RefRangeHigh: ptr(110), Aliases: []string{"CHLORIDE"}},
	{Name: "Carbon Dioxide", LOINCCode: "2028-9", Category: CategoryElectrolyte, Unit: "mmol/L", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(20), RefRangeHigh: ptr(32), Aliases: []string{"CARBON DIOXIDE", "CO2", "bicarbonate"}},
	{Name: "Calcium", LOINCCode: "17861-6", Category: CategoryMineral, Unit: "mg/dL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(8.6), RefRangeHigh: ptr(10.3), Aliases: []string{"CALCIUM"}},
	{Name: "Protein, Total", LOINCCode: "2885-2", Category: CategoryProtein, Unit: "g/dL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(6.1), RefRangeHigh: ptr(8.1), Aliases: []string{"PROTEIN, TOTAL", "total protein"}},
	{Name: "Albumin", LOINCCode: "1751-7", Category: CategoryProtein, Unit: "g/dL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(3.6), RefRangeHigh: ptr(5.1), PeptideRelevance: "liver synthetic function", Aliases: []string{"ALBUMIN"}},
	{Name: "Globulin", LOINCCode: "10834-0", Category: CategoryProtein, Unit: "g/dL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(1.9), RefRangeHigh: ptr(3.7), Aliases: []string{"GLOBULIN"}},
	{Name: "Bilirubin, Total", LOINCCode: "1975-2", Category: CategoryLiver, Unit: "mg/dL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(0.2), RefRangeHigh: ptr(1.2), Aliases: []string{"BILIRUBIN, TOTAL"}},
	{Name: "Alkaline Phosphatase", LOINCCode: "6768-6", Category: CategoryLiver, Unit: "U/L", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(36), RefRangeHigh: ptr(130), Aliases: []string{"ALKALINE PHOSPHATASE", "ALP"}},
	{Name: "AST", LOINCCode: "1920-8", Category: CategoryLiver, Unit: "U/L", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(10), RefRangeHigh: ptr(40), PeptideRelevance: "liver safety monitoring during peptide protocols", Aliases: []string{"AST", "SGOT", "aspartate aminotransferase"}},
	{Name: "ALT", LOINCCode: "1742-6", Category: CategoryLiver, Unit: "U/L", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(9), RefRangeHigh: ptr(46), PeptideRelevance: "liver safety monitoring during peptide protocols", Aliases: []string{"ALT", "SGPT", "alanine aminotransferase"}},

	// ── Hormonal Panel ──
	{Name: "Testosterone, Total", LOINCCode: "2986-8", Category: CategoryHormonal, Unit: "ng/dL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(250), RefRangeHigh: ptr(1100), PeptideRelevance: "GH secretagogues, peptide stacks, Gonadorelin", Aliases: []string{"TESTOSTERONE, TOTAL, MS", "total testosterone"}},
	{Name: "Testosterone, Free", LOINCCode: "2991-8", Category: CategoryHormonal, Unit: "pg/mL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(35.0), RefRangeHigh: ptr(155.0), PeptideRelevance: "bioavailable fraction, Gonadorelin protocol monitoring", Aliases: []string{"TESTOSTERONE, FREE", "free testosterone"}},
	{Name: "SHBG", LOINCCode: "13967-5", Category: CategoryHormonal, Unit: "nmol/L", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(10), RefRangeHigh: ptr(50), PeptideRelevance: "affects free testosterone availability", Aliases: []string{"SEX HORMONE BINDING GLOBULIN", "sex hormone binding globulin"}},
	{Name: "Estradiol", LOINCCode: "2243-4", Category: CategoryHormonal, Unit: "pg/mL", RefRangeType: RefRangeLessEqual, RefRangeHigh: ptr(39), PeptideRelevance: "aromatase activity during GH/anabolic protocols", Aliases: []string{"ESTRADIOL", "E2"}},
	{Name: "FSH", LOINCCode: "15067-2", Category: CategoryHormonal, Unit: "mIU/mL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(1.4), RefRangeHigh: ptr(12.8), PeptideRelevance: "pituitary function, Kisspeptin/Gonadorelin monitoring", Aliases: []string{"FSH", "follicle stimulating hormone"}},
	{Name: "LH", LOINCCode: "10501-5", Category: CategoryHormonal, Unit: "mIU/mL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(1.5), RefRangeHigh: ptr(9.3), PeptideRelevance: "pituitary function, Kisspeptin/Gonadorelin monitoring", Aliases: []string{"LH", "luteinizing hormone"}},
	{Name: "DHEA-S", LOINCCode: "2191-5", Category: CategoryHormonal, Unit: "mcg/dL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(61), RefRangeHigh: ptr(442), PeptideRelevance: "adrenal function, anti-aging marker", Aliases: []string{"DHEA SULFATE", "DHEA-S", "dehydroepiandrosterone sulfate"}},
	{Name: "Cortisol", LOINCCode: "2143-6", Category: CategoryHormonal, Unit: "mcg/dL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(4.0), RefRangeHigh: ptr(22.0), PeptideRelevance: "stress response, GHRP side effect (cortisol elevation)", Aliases: []string{"CORTISOL, TOTAL", "cortisol"}},
	{Name: "Insulin", LOINCCode: "2484-4", Category: CategoryMetabolic, Unit: "uIU/mL", RefRangeType: RefRangeTiered, OptimalHigh: ptr(18.4), PeptideRelevance: "insulin sensitivity, GH secretagogue metabolic impact", Aliases: []string{"INSULIN", "fasting insulin"}},

	// ── Thyroid ──
	{Name: "TSH", LOINCCode: "3016-3", Category: CategoryThyroid, Unit: "mIU/L", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(0.40), RefRangeHigh: ptr(4.50), PeptideRelevance: "thyroid function baseline", Aliases: []string{"TSH", "thyroid stimulating hormone"}},

	// ── Inflammatory ──
	{Name: "hs-CRP", LOINCCode: "30522-7", Category: CategoryInflammatory, Unit: "mg/L", RefRangeType: RefRangeTiered, OptimalHigh: ptr(1.0), PeptideRelevance: "systemic inflammation — BPC-157, TB-500, KPV efficacy marker", Aliases: []string{"HS CRP", "high-sensitivity C-reactive protein", "C-reactive protein"}},
	{Name: "Homocysteine", LOINCCode: "2028-9", Category: CategoryInflammatory, Unit: "umol/L", RefRangeType: RefRangeLessEqual, RefRangeHigh: ptr(13.5), PeptideRelevance: "cardiovascular risk, methylation status", Aliases: []string{"HOMOCYSTEINE"}},

	// ── CBC ──
	{Name: "WBC", LOINCCode: "6690-2", Category: CategoryCBC, Unit: "Thousand/uL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(3.8), RefRangeHigh: ptr(10.8), PeptideRelevance: "immune function, Thymosin Alpha-1 response", Aliases: []string{"WHITE BLOOD CELL COUNT", "white blood cells"}},
	{Name: "RBC", LOINCCode: "789-8", Category: CategoryCBC, Unit: "Million/uL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(4.20), RefRangeHigh: ptr(5.80), Aliases: []string{"RED BLOOD CELL COUNT", "red blood cells"}},
	{Name: "Hemoglobin", LOINCCode: "718-7", Category: CategoryCBC, Unit: "g/dL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(13.2), RefRangeHigh: ptr(17.1), Aliases: []string{"HEMOGLOBIN"}},
	{Name: "Hematocrit", LOINCCode: "4544-3", Category: CategoryCBC, Unit: "%", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(39.4), RefRangeHigh: ptr(51.1), Aliases: []string{"HEMATOCRIT"}},
	{Name: "Platelet Count", LOINCCode: "777-3", Category: CategoryCBC, Unit: "Thousand/uL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(140), RefRangeHigh: ptr(400), Aliases: []string{"PLATELET COUNT", "platelets"}},
	{Name: "Neutrophils, Absolute", LOINCCode: "751-8", Category: CategoryCBC, Unit: "cells/uL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(1500), RefRangeHigh: ptr(7800), Aliases: []string{"ABSOLUTE NEUTROPHILS"}},
	{Name: "Lymphocytes, Absolute", LOINCCode: "731-0", Category: CategoryCBC, Unit: "cells/uL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(850), RefRangeHigh: ptr(3900), PeptideRelevance: "adaptive immunity, Thymosin Alpha-1 response", Aliases: []string{"ABSOLUTE LYMPHOCYTES"}},

	// ── Metabolic / Diabetes ──
	{Name: "HbA1c", LOINCCode: "4548-4", Category: CategoryMetabolic, Unit: "%", RefRangeType: RefRangeLessThan, RefRangeHigh: ptr(5.7), PeptideRelevance: "long-term glucose control, GH secretagogue impact", Aliases: []string{"HEMOGLOBIN A1c", "hemoglobin A1c", "A1C", "glycated hemoglobin"}},

	// ── Vitamins ──
	{Name: "Vitamin D, 25-OH", LOINCCode: "1989-3", Category: CategoryVitamin, Unit: "ng/mL", RefRangeType: RefRangeMinMax, RefRangeLow: ptr(30), RefRangeHigh: ptr(100), PeptideRelevance: "immune function, bone density, recovery", Aliases: []string{"VITAMIN D,25-OH,TOTAL,IA", "vitamin D", "25-hydroxy vitamin D"}},
}
