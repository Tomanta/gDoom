# gDoom

gDoom is a Golang port of the Doom source code with added features.

## Current status

Very early development. Currently can read the following from a byte buffer (a WAD file):

- [x] header
- [x] lump directory entry
- [x] full lump directory
- [ ] level
    - [x] things
    - [x] linedefs
    - [x] sidedefs
    - [x] vertexes
    - [ ] segs
    - [ ] ssectors
    - [ ] nodes
    - [x] sectors
    - [ ] rejects
    - [ ] blockmap
- [ ] playpal
- [ ] colormap
- [ ] endoom
- [ ] texture1 and texture2
- [ ] pnames
- [ ] demo

## Next steps

After the above is complete, command line options to:

- [ ] load custom wads
- [ ] export map to disk