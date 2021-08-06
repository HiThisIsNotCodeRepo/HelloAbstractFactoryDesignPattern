# HelloAbstractFactoryDesignPattern
> [Source](https://golangbyexample.com/abstract-factory-design-pattern-go/)
## Core elements

1. Factory function ,all objects are generated from the factory function.

```
// Factory generation
adidasFactory, _ := getSportsFactory("adidas")
nikeFactory, _ := getSportsFactory("nike")
// Shoe generation
aNikeShoe := nikeFactory.makeShoe()
// Short generation
aNikeShort := nikeFactory.makeShort()
// Shoe generation
anAdidasShoe := adidasFactory.makeShoe()
// Short generation
anAdidasShort := adidasFactory.makeShort()
```

2. Interface ,all factory function return the interface.

```
type iSportsFactory interface {
	makeShoe() iShoe
	makeShort() iShort
}

func getSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	} else if brand == "nike" {
		return &nike{}, nil
	}
	return nil, fmt.Errorf("wrong brand type passed")
}

```