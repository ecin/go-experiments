package matrix

import (
  "fmt"
  "math"
)

/* A Matrix is either a scalar value, or a composition of
   four submatrices. Quadrants 0, 1, 2, 3 correspond to
   north west, north east, south west and south east.
*/
type Matrix struct {
  value float64
  level float64
  quadrants [4]*Matrix
}

func (m *Matrix) Set(x float64, y float64, value float64) {
  // Increase matrix level if necessary
  m.grow(math.Max(x, y))

  if x == 0 && y == 0 && m.Size() == 1 {
    m.value = value
  } else {
    nextQuadrant := m.nextQuadrant(x, y)
    x, y = m.translate(x, y, nextQuadrant)

    if m.quadrants[nextQuadrant] == nil {
      m.quadrants[nextQuadrant] = new(Matrix)
    }

    m.quadrants[nextQuadrant].Set(x, y, value)
  }
}

// Horribly named function.
func (m *Matrix) translate(x float64, y float64, quadrant int) (float64, float64) {
  var halfSize = m.Size() / 2

  switch quadrant {
    case 0:
      return x, y
    case 1:
      return x - halfSize, y
    case 2:
      return x, y - halfSize
    case 3:
      return x - halfSize, y - halfSize
    default:
      panic("WTF?!")
  }

  // Default return
  return 0, 0
}

func (m *Matrix) nextQuadrant(x float64, y float64) int {
  var quadrantLocation int
  var halfSize = m.Size() / 2

  switch {
  case x < halfSize && y < halfSize:
    quadrantLocation = 0
  case x >= halfSize && y < halfSize:
    quadrantLocation = 1
  case x < halfSize && y >= halfSize:
    quadrantLocation = 2
  case x >= halfSize && y >= halfSize:
    quadrantLocation = 3
  default:
    panic("WTF?!")
  }

  return quadrantLocation
}

func (m *Matrix) Size() float64 {
  return math.Pow(2, m.level)
}

func (m *Matrix) Get(x float64, y float64) interface{} {

  if x > m.Size() || y > m.Size() {
    // maybe return 0?
    return nil
  } else if x == 0 && y == 0 && m.Size() == 1 {
    return m.value
  } else {
    nextQuadrant := m.nextQuadrant(x, y)
    x, y = m.translate(x, y, nextQuadrant)

    return m.quadrants[nextQuadrant].Get(x, y)
  }

  // Default return
  return 0
}

func (m *Matrix) grow(size float64) float64 {
  for m.Size() <= size {
    newMatrix := new(Matrix)

    newMatrix.quadrants = m.quadrants
    newMatrix.value = m.value
    newMatrix.level = m.level


    m.quadrants = [4]*Matrix{}
    m.quadrants[0] = newMatrix
    m.level = m.level + 1
  }

  return m.level
}

func (m *Matrix) plus(n *Matrix) *Matrix {
  mn := new(Matrix)

  if m.isScalar() && n.isScalar() {
    mn.value = m.value + n.value
  } else {
    for i, _ := range mn.quadrants {
      mn.quadrants[i] = m.quadrants[i].plus(n.quadrants[i])
    }
  }

  return mn
}

func (m *Matrix) multiply(n *Matrix) *Matrix {
  mn := new(Matrix)

  // A Matrix times a zero Matrix is a zero Matrix
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

/* A Matrix is a zero Matrix iff:
   a) its value is zero
   b) all its quadrants are zero matrices
*/
func (m *Matrix) isZero() bool {
  if m == nil {
    return true
  }

  isZero := m.value == 0
  for _, quadrant := range m.quadrants {
    isZero = isZero && quadrant.isZero()
  }

  return isZero
}

/* A Matrix is a scalar iff:
   a) all its quadrants are nil
*/
func (m *Matrix) isScalar() bool {
  if m == nil {
    return false
  }

  isScalar := true
  for _, quadrant := range m.quadrants {
    isScalar = isScalar && !quadrant.isScalar()
  }

  return isScalar
}

func (m *Matrix) equals(n *Matrix) bool {
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

func (m *Matrix) String() string {
  return fmt.Sprintf("%i [%i %i %i %i]",
      m.value,
      m.quadrants[0].value,
      m.quadrants[1].value,
      m.quadrants[2].value,
      m.quadrants[3].value)
}
