package protocol

import "testing"

// TestMethodIDs checks the method table against the Ranking (SSB4) wiki page (IDs 14-20)
func TestMethodIDs(t *testing.T) {
	want := map[string]uint32{
		"GetCommunityCompetitionScore":        14,
		"GetCommunityCompetitionRankData":     15,
		"GetCommunityCompetitionRanking":      16,
		"PutCommunityCompetitionScore":        17,
		"DeleteCommunityCompetitionScore":     18,
		"DeleteCommunityCompetitionScores":    19,
		"DEBUG_PutCommunityCompetitionScores": 20,
	}
	got := map[string]uint32{
		"GetCommunityCompetitionScore":        MethodGetCommunityCompetitionScore,
		"GetCommunityCompetitionRankData":     MethodGetCommunityCompetitionRankData,
		"GetCommunityCompetitionRanking":      MethodGetCommunityCompetitionRanking,
		"PutCommunityCompetitionScore":        MethodPutCommunityCompetitionScore,
		"DeleteCommunityCompetitionScore":     MethodDeleteCommunityCompetitionScore,
		"DeleteCommunityCompetitionScores":    MethodDeleteCommunityCompetitionScores,
		"DEBUG_PutCommunityCompetitionScores": MethodDEBUG_PutCommunityCompetitionScores,
	}

	for name, id := range want {
		if got[name] != id {
			t.Errorf("%s = %d, want %d", name, got[name], id)
		}
	}

	if ProtocolID != 112 {
		t.Errorf("ProtocolID = %d, want 112", ProtocolID)
	}

	if len(patchedMethods) != 7 {
		t.Fatalf("patchedMethods has %d entries, want 7", len(patchedMethods))
	}

	seen := map[uint32]bool{}
	for _, id := range patchedMethods {
		if id < 14 || id > 20 || seen[id] {
			t.Errorf("unexpected or duplicate patched method %d", id)
		}
		seen[id] = true
	}
}
