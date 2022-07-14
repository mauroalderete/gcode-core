// gcode package contains the model to represent a single gcode expression.
//
// Gcode expressions are the minimal pieces in a block (line with a command at the gcode file).
//
// A gcode can be interpreted as a command, parameter, or for any other special purpose.
//
// It consists of only a stand-alone letter, named word.
// Or the word directly followed by a numeric or alphanumeric value named address.
// The word gives information about the meaning of the gcode.
// Instead, the address asserts a value to gcode.
// It can be integers (128) or fractional numbers (12.42) or strings depending on the context.
//
// For example, an X coordinate can take integers (X175) or fractionals (X17.62).
//
// gcode package implements two constructor methods to instantiate a Gcode struct or a GcodeAddressable struct.
//
// One of them is used to model a gcode with a stand-alone word.
// Is said, allows to get a gcode without an address component.
//
// The other method is a generic method that allows constructing a typical gcode object but, also,
// includes an address struct of the data type defined by the AddressType interface in the address package.
// This constructor we allow to get a gcode with a string address, or address int32 or address float32.
//
// All gcode types implement a Gcoder interface.
// Instead, the gcodes that contain an address implement the GcodeAddresser interface.
//
// GcodeAddresser interface wrap Gcoder interface.
// This means, that all the GcodeAddresser objects are as well Gcoder objects.
//
// The word model represents the letter of a gcode command that identifies a single action or a parameter.
//
// The words consist of a simple struct that implements a single value string. This value is a letter according to specification to gcode.
// Each word contains a single letter in uppercase. Each one of them can come accompanied by an address or not,
// or it needs other gcodes that give more information in the form of parameters.
//
// Exist, different classes the words, some are commands and others are used likes parameters for the commands.
//
// The address model allows store and management the representation of the part assigned to the value of a gcode.
//
// A gcode can have or doesn't have an address. When it has, the address must be of either int32, float32 or string data type.
// This package contains a constructor that returns an address of some of these data types defined by the AddressType interface.
//
// An address struct is bound with a series of methods and functions that allow you to operate with the value of the address.
package gcode
