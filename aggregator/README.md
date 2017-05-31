Feed Aggregation
================

What's done here:

+ Make async requests to given URLs
+ Parse RSS feed into a struct
+ Write feed items into file or STDOUT through buffer
+ Wait for all async go-routines to end

HOW TO USE
----------

```
% make
% ./aggregator [-output path/to/output/file.txt]
```

+ `-output`: (Optional)
