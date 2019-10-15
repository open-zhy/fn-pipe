# fn-pipe

[![CircleCI](https://circleci.com/gh/open-zhy/fn-pipe.svg?style=svg)](https://circleci.com/gh/open-zhy/fn-pipe)

Functions pipeline for golang projects

## Basic Usage
```go
    // create the pipeline
    p, err := NewPipeline(
    	func(a int) (int, int) { return a + 5, a - 5 }, // fn1
    	func(a int, b int) int { return a * b },        // fn2
    )
    
    // error checking
    if err != nil {
    	fmt.Fatalf("Error while creating pipeline -> %s", err)
    }
    
    // output the result which should be fn2(fn1(...args))
    err, res := p.ExecWith(0) // res = [-25]
```


## Contributing
Please follow this [link](https://github.com/open-zhy/fn-pipe/blob/master/CONTRIBUTING.md)