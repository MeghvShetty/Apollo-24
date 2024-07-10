package logic

type TriageInput struct {
	DataClassification      string `form:DataClassification" binding:"required"`
	MultipleReleaseOveryear string `form:MultipleReleaseOveryear" binding:"required"`
	IntegrityRating         string `form:IntegrityRating" binding:"required"`
	NewInfrastructure       string `form:NewInfrastructure" binding:"required"`
	NewConnectivity         string `form:NewConnectivity" binding:"required"`
	RegulatoryRequirement   string `form:RegulatoryRequirement" binding:"required"`
	UserImpact              string `form:UserImpact" binding:"required"`
	PatternUsed             string `form:PatterUsed" binding:"required"`
}

func SirCalculator(input *TriageInput) string {
	var sir string = "SIR0"

	if ((input.DataClassification == "Highly Confidential" || input.DataClassification == "Confidential") && (input.IntegrityRating == "Severe" || input.IntegrityRating == "High") && input.MultipleReleaseOveryear == "Yes") && (input.PatternUsed == "Hybrid" ||
		input.PatternUsed == "New") {
		sir = "SIR3"
	} else if (input.NewInfrastructure == "Yes" || input.NewConnectivity == "Yes" && input.MultipleReleaseOveryear == "Yes") && (input.PatternUsed == "Hybrid" || input.PatternUsed == "New") {
		sir = "SIR2"
	} else if (input.NewInfrastructure == "Yes" || input.NewConnectivity == "Yes" || input.RegulatoryRequirement == "Yes" || input.UserImpact == ">10,000") && (input.PatternUsed == "Hybrid" || input.PatternUsed == "New") {
		sir = "SIR1"
	} else {
		sir = "SIR0"
	}

	return sir
}
