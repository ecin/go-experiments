package matrix

import "fmt"

/* A matrix is either a scalar value, or a composition of
   four submatrices. Quadrants 0, 1, 2, 3 correspond to
   north west, north east, south west and south east.
*/
type matrix struct {
  value int
  quadrants [4]*matrix
}

func (m *matrix) plus(n *matrix) *matrix {
  mn := new(matrix)

  if m.isScalar() && n.isScalar() {
    mn.value = m.value + n.value
  } else {
    for i, _ := range mn.quadrants {
      mn.quadrants[i] = m.quadrants[i].plus(n.quadrants[i])
    }
  }

  return mn
}

func (m *matrix) multiply(n *matrix) *matrix {
  mn := new(matrix)

  // A matrix times a zero matrix is a zero matrix
  if m.isZero() || n.isZero() {
    return mn
  }

  if m.isScalar() && n.isScalar() {
    mn.value = m.value * n.value
  } else {
    mn.quadrants[0] = m.quadrants[0].multiply(n.quadrants[0]).plus(m.quadrants[1].multiply(n.quadrants[2]))
    mn.quadrants[1] = m.quadrants[0].multiply(n.quadrants[1]).plus(m.quadrants[1].multiply(n.quadrants[3]))
    mn.quadrants[2] = m.quadrants[2].multiply(n.quadrants[0]).plus(m.quadrants[3].multiply(n.quadrants[2]))
    mn.quadrants[3] = m.quadrants[2].multiply(n.quadrants[1]).plus(m.quadrants[3].multiply(n.quadrants[3]))
  }

  return mn
}

/* A matrix is a zero matrix iff:
   a) its value is zero
   b) all its quadrants are zero matrices
*/
func (m *matrix) isZero() bool {
  if m == nil {
    return true
  }

  isZero := m.value == 0
  for _, quadrant := range m.quadrants {
    isZero = isZero && quadrant.isZero()
  }

  return isZero
}

/* A matrix is a scalar iff:
   a) all its quadrants are nil
*/
func (m *matrix) isScalar() bool {
  if m == nil {
    return false
  }

  isScalar := true
  for _, quadrant := range m.quadrants {
    isScalar = isScalar && !quadrant.isScalar()
  }

  return isScalar
}

func (m *matrix) equals(n *matrix) bool {
  if m == nil && n == nil {
    return true
  }

  isEqual := true
  isEqual = isEqual && m.value == n.value
  for i, _ := range m.quadrants {
    isEqual = isEqual && m.quadrants[i].equals(n.quadrants[i])
  }

  return isEqual
}

func (m *matrix) String() string {
  return fmt.Sprintf("%i [%i %i %i %i]",
      m.value,
      m.quadrants[0].value,
      m.quadrants[1].value,
      m.quadrants[2].value,
      m.quadrants[3].value)
}
