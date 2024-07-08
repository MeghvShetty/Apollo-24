package logic

type TriageInput struct {
	DataClassification      string
	MultipleReleaseOveryear string
	IntegrityRating         string
	NewInfrastructure       string
	NewConnectivity         string
	RegulatoryRequirement   string
	UserImpact              string
	PatternUsed             string
}

func SirCalculator(input *TriageInput) string {
	var sir string = "0"

	if ((input.DataClassification == "Highly Confidential" || input.DataClassification == "Confidential") && (input.IntegrityRating == "Severe" || input.IntegrityRating == "High") && input.MultipleReleaseOveryear == "Yes") && (input.PatternUsed == "Hybrid" ||
		input.PatternUsed == "New") {
		sir = "3"
	} else if (input.NewInfrastructure == "Yes" || input.NewConnectivity == "Yes" && input.MultipleReleaseOveryear == "Yes") && (input.PatternUsed == "Hybrid" || input.PatternUsed == "New") {
		sir = "2"
	} else if (input.NewInfrastructure == "Yes" || input.NewConnectivity == "Yes" || input.RegulatoryRequirement == "Yes" || input.UserImpact == ">10,000") && (input.PatternUsed == "Hybrid" || input.PatternUsed == "New") {
		sir = "1"
	} else {
		sir = "0"
	}

	return sir
}
