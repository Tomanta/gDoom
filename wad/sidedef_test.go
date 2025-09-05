package wad

/*
func TestNewSidedefssFromBytes(t *testing.T) {
	t.Run("returns error if buffer wrong length", func(t *testing.T) {
		data := []byte{
			0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF,
			0x01, 0x00, 0x02, 0x00, 0x01, 0x00, 0x00,
		}
		var numSidedefs int32 = 2
		_, err := NewSidedefsFromBytes(data, numSidedefs)
		if err == nil {
			t.Fatalf("did not receive expected error")
		}
	})

	t.Run("returns correct information", func(t *testing.T) {
		data := []byte{
			0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF,
			0x01, 0x00, 0x02, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0xFF, 0xFF,
		}
		var numSideedefs int32 = 2
		want := []Sidedef{
			{
	TextureOffsetX: ,
	TextureOffsetY: ,
	SectorID      : ,
	UpperTexture  : ,
	LowerTexture  : ,
	MiddleTexture : ,
			}, {
	TextureOffsetX: ,
	TextureOffsetY: ,
	SectorID      : ,
	UpperTexture  : ,
	LowerTexture  : ,
	MiddleTexture : ,
			},
		}
		got, err := NewSidedefsFromBytes(data, numSidedefs)
		if err != nil {
			t.Fatalf("could not read sidedefs: %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v, got %v", want, got)
		}

	})
}
*/
