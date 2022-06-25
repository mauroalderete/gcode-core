# CHANGELOG

# 25/06/2022

## Feature

I reimplement package structs and, I create new structs to manage each of the elements that compose a gcode' line.

A block will be composed of multiple gcode' struct that represent each part of the command line, with exception of the comment section.

I do not implement a system to validate a block using some map of allowed commands yet.

Check struct consist in a interface that is implement for diferents algorithms of verification like checksum or crc.

In fact, crc algorithm is not implemented ready, solo exists as a dummy sketch.

Address package allows will construction of an address struct from multiple data types.

Same time, block struct expose different format of representation that include or ommite comments, check or data.

## Plans

I will complete each package with the respective documentation using gdoc system, later.
