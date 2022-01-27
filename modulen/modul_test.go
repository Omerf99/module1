package modul

import "testing"

func TestInitModul(t *testing.T) {
    wanted := false; // Bruker semantikken i Golangs konseptuelle modell
    state := TestInitModul();
    if state != wanted {
    t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
    }
}
