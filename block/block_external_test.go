package block_test

////#region Mocks

// type mockGcodeFactory struct{}

// // NewGcode is the constructor to instance a Gcode struct.
// //
// // Receive a word that represents the letter of the command of a gcode.
// //
// // If the word is an unknown symbol it returns nil with an error description.
// func (g *mockGcodeFactory) NewUnaddressableGcode(word byte) (gcode.Gcoder, error) {
// 	ng, err := unaddressablegcode.New(word)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ng, nil
// }

// func (g *mockGcodeFactory) NewAddressableGcodeUint32(word byte, address uint32) (gcode.AddresableGcoder[uint32], error) {

// 	ng, err := addressablegcode.New(word, address)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ng, nil
// }

// func (g *mockGcodeFactory) NewAddressableGcodeInt32(word byte, address int32) (gcode.AddresableGcoder[int32], error) {

// 	ng, err := addressablegcode.New(word, address)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ng, nil
// }

// func (g *mockGcodeFactory) NewAddressableGcodeFloat32(word byte, address float32) (gcode.AddresableGcoder[float32], error) {

// 	ng, err := addressablegcode.New(word, address)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ng, nil
// }

// func (g *mockGcodeFactory) NewAddressableGcodeString(word byte, address string) (gcode.AddresableGcoder[string], error) {

// 	ng, err := addressablegcode.New(word, address)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ng, nil
// }

// //#endregion

// func TestNew(t *testing.T) {

// 	t.Run("prueba1", func(t *testing.T) {

// 		// create a simple unaddressable gcode that will be the main command of the new block
// 		gc, err := unaddressablegcode.New('M')
// 		if err != nil {
// 			t.Errorf("got error %v, want error nil", err)
// 		}
// 		// invoke New block package function and adding two options callbacks to configure him
// 		b, err := block.New(gc, func(bc block.BlockConfigurer) error {

// 			// first option callback configures the minimal options
// 			// adding a checksum that implement hash.Hash, used to calculate each checksum of the block
// 			c := checksum.New()
// 			bc.Checksum(c)

// 			// adding a mock of GcoderFactory, used internally by block to handle each gcode that its composed.
// 			gcf := &mockGcodeFactory{}
// 			bc.GcodeFactory(gcf)

// 			// adding a line number gcode to indexer this new block
// 			gcn, err := addressablegcode.New[uint32]('N', 100)
// 			if err != nil {
// 				t.Errorf("failed to create lineNumber gcode: got error %v, want error nil", err)
// 			}

// 			fmt.Println("linenumber: ", gcn)
// 			bc.LineNumber(gcn)
// 			return nil

// 		}, func(bc block.BlockConfigurer) error {
// 			// the second option callback try used a new configuration option to set the slice the gcode parameters

// 			fmt.Println("agrego parametros??")
// 			// For this, we verify by assertion if the bc instance implement the new BlockConfigurerParameters interface
// 			if bcp, ok := bc.(block.BlockConfigurerParameters); ok {
// 				// If is true then we can invoke the Parameters method to adding a slice of Gcoder
// 				fmt.Println("yes! agrego parametros")
// 				gc1, err := unaddressablegcode.New('X')
// 				if err != nil {
// 					log.Fatal("Oops!")
// 				}
// 				gc2, err := unaddressablegcode.New('Y')
// 				if err != nil {
// 					log.Fatal("Oops!")
// 				}
// 				gc3, err := unaddressablegcode.New('Z')
// 				if err != nil {
// 					log.Fatal("Oops!")
// 				}
// 				bcp.Parameters([]gcode.Gcoder{gc1, gc2, gc3})
// 			} else {
// 				return fmt.Errorf("Ups, failed loading parameters at new block: BlockConfigurer not implement BlockConfigurerParameters yet")
// 			}
// 			return nil
// 		})

// 		if err != nil {
// 			log.Fatal("Oops! x2")
// 		}

// 		fmt.Println("block: ", b.ToLineWithCheckAndComments())
// 	})
// }
