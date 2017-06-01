Feed Aggregation in Parallel
============================

What's done here:

+ Boot 4 workers
+ Send feed URLs to workers
+ Workers fetch RSS feed, parse RSS and write to file
+ Continue this until all URLs are processed
+ Finish batch once all done

HOW TO USE
----------

```
% make
% ./aggregator [-output path/to/output/file.txt]
```

+ `-output`: (Optional)
