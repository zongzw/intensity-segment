# Intensity Segments

## Repository Description

This repository presents IntensitySegements problem and solution in golang and javascript.

Find the details in the separate folders.

## Problem Description

We are looking for a program that manages “intensity” by segments. 

Segments are intervals from -infinity to infinity, we’d like you to implement functions that updates intensity by an integer amount for a given range. 

All intensity starts with 0. Please implement these three functions:

```javascript
/** *
 * NOTE: Feel free to add any extra member variables/functions you like.
 */
export class IntensitySegments {
    /**
     *
     */
    add(from, to, amount) {
        // TODO: implement this
    }
    /** 
     *
     */
    set(from, to, amount) {
        // TODO: implement this
    }
    /** 
     *
     */
    toString() {
        // TODO: implement this
    }
}


// Here is an example sequence:
// (data stored as an array of start point and value for each segment.) 

const segments = new IntensitySegments();
segments.toString(); // Should be "[]"
segments.add(10, 30, 1);
segments.toString(); // Should be: "[[10,1],[30,0]]"
segments.add(20, 40, 1);
segments.toString(); // Should be: "[[10,1],[20,2],[30,1],[40,0]]"
segments.add(10, 40, -2);
segments.toString(); // Should be: "[[10,-1],[20,0],[30,-1],[40,0]]"

// Another example sequence:
const segments = new IntensitySegments(); 
segments.toString(); // Should be "[]"
segments.add(10, 30, 1);
segments.toString(); // Should be "[[10,1],[30,0]]"
segments.add(20, 40, 1);
segments.toString(); // Should be "[[10,1],[20,2],[30,1],[40,0]]"
segments.add(10, 40, -1);
segments.toString(); // Should be "[[20,1],[30,0]]"
segments.add(10, 40, -1);
segments.toString(); // Should be "[[10,-1],[20,0],[30,-1],[40,0]]"

```