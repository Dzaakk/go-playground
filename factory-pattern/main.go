package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func (t Triangle) Area() float64 {
	s := (t.SideA + t.SideB + t.SideC) / 2
	return math.Sqrt(s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC))
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

type ShapeFactory struct{}

func (f *ShapeFactory) CreateShape(shapeType string, params ...float64) (Shape, error) {
	switch shapeType {
	case "circle":
		if len(params) != 1 {
			return nil, fmt.Errorf("circle requires 1 parameter (radius)")
		}
		return Circle{Radius: params[0]}, nil
	case "rectangle":
		if len(params) != 2 {
			return nil, fmt.Errorf("rectangle requires 2 parameters (width, height)")
		}
		return Rectangle{Width: params[0], Height: params[1]}, nil
	case "triangle":
		if len(params) != 3 {
			return nil, fmt.Errorf("triangle requires 3 parameters (side a, side b, side c)")
		}
		return Triangle{SideA: params[0], SideB: params[1], SideC: params[2]}, nil

	default:
		return nil, fmt.Errorf("unkonw shape type: %s", shapeType)
	}
}

type ShapeOption func(Shape) Shape

func WithMinimumSize(minArea float64) ShapeOption {
	return func(s Shape) Shape {
		if s.Area() < minArea {
			fmt.Printf("Warning: Shape area (%.2f) is less than minum (%.2f)\n", s.Area(), minArea)
		}
		return s
	}
}

type ConfigurableShapeFactory struct{}

func (f *ConfigurableShapeFactory) CreateShape(shapeType string, options ...ShapeOption) (Shape, error) {
	var shape Shape
	var err error

	switch shapeType {
	case "Circle":
		shape = Circle{Radius: 1.0}
	case "rectangle":
		shape = Rectangle{Width: 1.0, Height: 1.0}
	case "triangle":
		shape = Triangle{SideA: 1.0, SideB: 1.0, SideC: 1.0}
	default:
		return nil, fmt.Errorf("unkown shape type: %s", shapeType)
	}

	for _, option := range options {
		shape = option(shape)
	}

	return shape, err
}

func ExampleFactoryUsage() {

	basicFactory := &ShapeFactory{}

	circle, err := basicFactory.CreateShape("circle", 5.0)
	if err != nil {
		fmt.Println("Error creating circle:", err)
		return
	}
	fmt.Printf("Circle - Area: %.2f, Perimeter: %.2f\n", circle.Area(), circle.Perimeter())

	rectangle, err := basicFactory.CreateShape("rectangle", 4.0, 5.0)
	if err != nil {
		fmt.Println("Error creating rectangle:", err)
		return
	}
	fmt.Printf("Circle - Area: %.2f, Perimeter: %.2f\n", rectangle.Area(), rectangle.Perimeter())

	configurableFactory := &ConfigurableShapeFactory{}

	configuredCircle, err := configurableFactory.CreateShape(
		"circle",
		WithMinimumSize(50.0),
	)
	if err != nil {
		fmt.Println("Error creating configured circle: ", err)
		return
	}
	fmt.Printf("Configured Circle - Area: %.2f\n", configuredCircle.Area())
}
