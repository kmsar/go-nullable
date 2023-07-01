#Nullable Types in GO
A nullable type can represent the correct range of values for its underlying value type, plus an additional null value
A nullable value type T represents all values of its underlying value type T and an additional null value.
For example, you can assign any of the following three values to a Int (not int) variable: int, or null. 
An underlying value type T cannot be a nullable value type itself.

    func TestInt(t *testing.T) {
        a := NewNullable[int]()
        if !a.HasValue() {
            //vale is null
        }
    
        a.SetValue(1)
        if !a.HasValue() {
            t.Fail()
        }
        println(a.GetValue())
    
        b := NewNullable[int]()
        println(b.GetValueOr(2))
    }

