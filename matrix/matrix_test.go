package matrix

import (
  "testing"
)

func TestSetGet(t *testing.T) {
  m := new(Matrix)

  var examples = []struct {
    x float64
    y float64
    value float64
  }{
    {0, 0, 0},
    {1, 1, 1},
    {2, 1, 1},
    {64, 64, -1},
    {1024, 1024, 0},
    {3, 3, 3},
  }

  for _, example := range examples {
    m.Set(example.x, example.y, example.value)

    var value = m.Get(example.x, example.y)
    if (value != example.value) {
      t.Error("Expected value at (%d, %d) to be %d, was %d",
        example.x,
        example.y,
        example.value,
        value)
    }
  }
}

func TestIsZero(t *testing.T) {
  m := new(matrix)

  if (!m.isZero()) {
    t.Error("A new matrix should be zero")
  }

  m.value = 5

  if (m.isZero()) {
    t.Error("A matrix with a value shouldn't be zero")
  }

  m = new(matrix)
  m2 := new(matrix)
  m.quadrants[0] = m2

  if (!m.isZero()) {
    t.Error("A matrix with a zero quadrant should be zero")
  }

  m2.value = 1
  if (m.isZero()) {
    t.Error("A matrix with a non-zero quadrant shouldn't be zero")
  }
}

func TestIsScalar(t *testing.T) {
  m := new(matrix)

  if (!m.isScalar()) {
    t.Error("A new matrix should be a scalar")
  }

  m.quadrants[0] = new(matrix)

  if (m.isScalar()) {
    t.Error("A matrix with a quadrant shouldn't be a scalar")
  }

}

func Testequals(t *testing.T) {
  m := new(matrix)
  n := new(matrix)

  if (!m.equals(n)) {
    t.Error("Two new matrices should be equal")
  }

  m.value = 1

  if (m.equals(n)) {
    t.Error("Matrice with different values should not be equal")
  }

  n.value = m.value

  if (!m.equals(n)) {
    t.Error("Matrices with equal values (and empty quadrants) should be equal")
  }

  o := new(matrix)
  p := new(matrix)
  for i, _ := range o.quadrants {
    o.quadrants[i] = new(matrix)
    o.quadrants[i].value = i

    p.quadrants[i] = new(matrix)
    p.quadrants[i].value = i
  }

  if (!o.equals(p)) {
    t.Error("Matrices with equal values and equal quadrants should be equal")
  }

  o.quadrants[0].value = 100

  if (m.equals(n)) {
    t.Error("Matrices with equal value but different quadrants shouldn't be equal")
  }

}

func TestPlus(t *testing.T) {
  m := new(matrix)
  n := new(matrix)

  o := m.plus(n)

  if (o.value != 0) {
    t.Error("The sum of two new matrices should have a value of 0")
  }

  for i, _ := range m.quadrants {
    m.quadrants[i] = new(matrix)
    m.quadrants[i].value = i

    n.quadrants[i] = new(matrix)
    n.quadrants[i].value = i
  }

  o = m.plus(n)

  if (o.isScalar()) {
    t.Error("The sum of two non-scalar matrices should not be a scalar")
  }


  for i, quadrant := range o.quadrants {
    if (quadrant.value != i + i) {
      t.Error("The sum of two non-scalar matrices should add its quadrants")
    }
  }

}

func TestMultiply(t *testing.T) {
  m := new(matrix)
  n := new(matrix)

  m.value = 3
  n.value = 5

  o := m.multiply(n)

  if (o.value != m.value * n.value) {
    t.Error("The product of two scalar matrices should be the product of their values")
  }

  for i, _ := range m.quadrants {
    m.quadrants[i] = new(matrix)
    m.quadrants[i].value = 1

    n.quadrants[i] = new(matrix)
    n.quadrants[i].value = 1
  }

  o = m.multiply(n)
  p := new(matrix)

  for i, _ := range p.quadrants {
    p.quadrants[i] = new(matrix)
    p.quadrants[i].value = 2
  }

  if (!o.equals(p)) {
    t.Error("The product of two matrices full of ones should be a matrix full of twos")
  }

}
