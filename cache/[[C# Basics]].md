# C# Basics

# 🥝C#

[https://docs.microsoft.com/en-us/dotnet/csharp/tour-of-csharp/](https://docs.microsoft.com/en-us/dotnet/csharp/tour-of-csharp/)

## 1. Basic program structure

    
    using System;
    
    namespace Azure{
    	class AzureService{
    		static public void Main(string[] args){}
    	}
    }

## 2. DataType

### Base Types

RendeRenderCollectionViewrCode NYI
- `sizeof` : `sizeof(int) //4`

### Reference Type

**object = System.Object**

- ultimate base class: can be assigned values of any type
- boxing: value type ⇒ object type
- unboxing: object type ⇒ value type

**dynamic**

- type checking for **object type** variables takes place at **compile time**
- type checking for **dynamic type** variables takes place at **run time**

**string = System.String**

> @quoted

### Pointer Type

- `type* identifier`

## 3. Type conversion

**implicit type conversion**

- from derived classes to base classes
- from smaller to larger integral type

**explicit type conversion**

- cast operator

**Built-in Conversion Methods**

- e.g. `ToBoolean` , `ToByte` ...

## 4. Console Interaction

- `Console.Readline`
- `Console.Writeline`

## 5. Operations

### Miscellaneous operations

- `sizeof`
- `typeof`

## 6. Encapsulation

**Access Specifier**

- Public
- Private: class scope
- Protected: child class scope
- Internal: application scope
- Protected Internal:

## 7. Nullable Identifiers

- ?: `<data_type> ? <variable_name> = null;`
- ??: `num2 = nullable_num ?? 10 // 10`

## 8. Array

> [https://www.tutorialspoint.com/csharp/csharp_arrays.htm](https://www.tutorialspoint.com/csharp/csharp_arrays.htm)

## 9. Structure

- methods ✔️
- default constructor ✖️

## 10. Enum

    enum ENUM_NAME{
    	enumeration list
    }

## 11. Polymorphism

> Overwritten

## 12. Operator Method

    public static Box operator+ (Box b, Box c) {
       Box box = new Box();
       box.length = b.length + c.length;
       box.breadth = b.breadth + c.breadth;
       box.height = b.height + c.height;
       return box;
    }
    
    // application
    Box a = new Box(1,2,3);
    Box b = new Box(5,6,7);
    a + b // Box(6,8,10)

## 13. Interface

> syntactical contract

    namespace InterfaceApplication {
    	public interface ITransaction {
    		void showTransaction();
    	}
    
    	public class StockTransaction : ITransaction {
    		public void showTransaction(){ ... }
    	}
    
    	public class FundTransaction {
    		public void showTransaction(){ ... }
    	}
    }

## 14. Namespace

- namespace `namespace <namespace_name>`
- using

    `using System` means to import variables from System namespace.

    which makes `Console.WriteLine` equal to `System.Console.WriteLine`

Namespaces can be nested.

    namespace father {
    	namespace child {
    	}
    }

## 15. 🔺Preprocessor

RendeRenderCollectionViewrCode NYI

## 16. Exception handler

**Syntax**

    try {
       // statements causing exception
    } catch( ExceptionName e1 ) {
       // error handling code
    } catch( ExceptionName e2 ) {
       // error handling code
    } catch( ExceptionName eN ) {
       // error handling code
    	 throw(new System.IO.IOException("Raise IO Exception"))
    } finally {
       // statements to be executed
    }

## 17. I/O

**stream**

- input stream: handling writting
- output stream: handling reading

# C# MS docs

[https://docs.microsoft.com/en-us/dotnet/csharp/tour-of-csharp/](https://docs.microsoft.com/en-us/dotnet/csharp/tour-of-csharp/)

    (double Sum, int Count) t2 = (4.5, 3);
    Console.WriteLine($"Sum of {t2.Count} elements is {t2.Sum}.");
    // Output:
    // Sum of 3 elements is 4.5.

field ⇒ 变量 variable

- static field
- instance field

    public override string ToString() => "This is an object";

## Parameter types of method

**Ref parameter**

    static void Swap(ref int x, ref int y)
    {
    	int temp = x;
    	x = y;
    	y = temp;
    }

**Output parameter**

    static void Divide(int x, int y, out int result, out int remainder)
    {
        result = x / y;
        remainder = x % y;
    }
    
    public static void OutUsage()
    {
        Divide(10, 3, out int res, out int rem);
        Console.WriteLine($"{res} {rem}");	// "3 1"
    }

**Parameter array**

    public class Console
    {
        public static void Write(string fmt, params object[] args) { }
        public static void WriteLine(string fmt, params object[] args) { }
        // ...
    }

**Constructor**

- static constructor: first intro classjie
- instance constructor

**Finalizer**

**Property**

    public class PropertyClass
    {
    	private int _volumn = 100;
    	public int Volumn
    	{
    		get => _volumn;
    		set
    		{
    			if (value < 0)
    			{
    				_volumn = 0;
    			}
    			else if (value > 100)
    			{
    				_volumn = 100;	
    			}
    			else
    			{
    				_volumn = value;
    			}
    		}
    	}
    }

**Indexers**

    public T this[int index]
    {
        get => _items[index];
        set
        {
            _items[index] = value;
            OnChanged();
        }
    }

**Events**

    class EventTest
    {
    	private int _val;
    	public int val
    	{
    		set
    		{
    			_val = value;
    			onChanged();
    		}
    		get
    		{
    			return _val;
    		}
    	}
    
    	protect virtual void OnChanged() 
    	{
    		Changed?.invoke(this, EventArgs.Empty)
    	}
    
    	public event EventHandler Changed
    	{
    		add
    		{
    		}
    		
    		remove
    		{
    		}
    	}
    }
    
    // Main
    EventTest test = new EventTest();
    test.Changed += new EventHandler( () => {} )

**Operators**

    // class scope
    public static bool operator == (T a, T b) => a.val == b.val;

**Finalizer**

    class Car
    {
    	Car() {} // constructor
    	~Car() {} // deconstructor
    }

## `virtual` vs. `abstract`

    public abstract class E
    {
        public abstract void AbstractMethod(int i);
    
        public virtual void VirtualMethod(int i)
        {
            // Default implementation which can be overridden by subclasses.
        }
    }
    
    public class D : E
    {
        public override void AbstractMethod(int i)
        {
            // You HAVE to override this method
        }
        public override void VirtualMethod(int i)
        {
            // You are allowed to override this method.
        }
    }

## "⇒" in C#

[=> operator - C# reference](https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/operators/lambda-operator)

## Statements

- A *block* permits multiple statements to be written in contexts where a single statement is allowed. A block consists of a list of statements written between the delimiters `{` and `}`.
- *Declaration statements* are used to declare local variables and constants.
- *Expression statements* are used to evaluate expressions. Expressions that can be used as statements include method invocations, object allocations using the `new` operator, assignments using `=` and the compound assignment operators, increment and decrement operations using the `++` and `-` operators and `await` expressions.
- *Selection statements* are used to select one of a number of possible statements for execution based on the value of some expression. This group contains the `if` and `switch` statements.
- *Iteration statements* are used to execute repeatedly an embedded statement. This group contains the `while`, `do`, `for`, and `foreach` statements.
- *Jump statements* are used to transfer control. This group contains the `break`, `continue`, `goto`, `throw`, `return`, and `yield` statements.
- The `try`...`catch` statement is used to catch exceptions that occur during execution of a block, and the `try`...`finally` statement is used to specify finalization code that is always executed, whether an exception occurred or not.
- The `checked` and `unchecked` statements are used to control the overflow-checking context for integral-type arithmetic operations and conversions.
- The `lock` statement is used to obtain the mutual-exclusion lock for a given object, execute a statement, and then release the lock.
- The `using` statement is used to obtain a resource, execute a statement, and then dispose of that resource.

## LINQ: Language-Integrated Query

[https://docs.microsoft.com/en-us/dotnet/api/system.linq?view=net-5.0](https://docs.microsoft.com/en-us/dotnet/api/system.linq?view=net-5.0)

[https://github.com/dotnet/try-samples/blob/main/101-linq-samples](https://github.com/dotnet/try-samples/blob/main/101-linq-samples)

## Delegate

    public delegate int PerformCalculation(int x, int y);

## Convert and Parse

- `Int64.parse(132)`
- `Convert.ToInt64(123)`

## `Record` vs `Class`

[Records](https://docs.microsoft.com/en-us/dotnet/csharp/fundamentals/types/records)

`public record Person(string FirstName, string LastName);`

    public record Person(string FirstName, string LastName)
    {
        public string[] PhoneNumbers { get; init; }
    }
    
    public static void Main()
    {
        Person person1 = new("Nancy", "Davolio") { PhoneNumbers = new string[1] };
        Console.WriteLine(person1);
        // output: Person { FirstName = Nancy, LastName = Davolio, PhoneNumbers = System.String[] }
    
        Person person2 = person1 with { FirstName = "John" };
        Console.WriteLine(person2);
        // output: Person { FirstName = John, LastName = Davolio, PhoneNumbers = System.String[] }
        Console.WriteLine(person1 == person2); // output: False
    
        person2 = person1 with { PhoneNumbers = new string[1] };
        Console.WriteLine(person2);
        // output: Person { FirstName = Nancy, LastName = Davolio, PhoneNumbers = System.String[] }
        Console.WriteLine(person1 == person2); // output: False
    
        person2 = person1 with { };
        Console.WriteLine(person1 == person2); // output: True
    }

## Anonymous Type

[Anonymous Types](https://docs.microsoft.com/en-us/dotnet/csharp/fundamentals/types/anonymous-types)

## Related tools

OmniSharp: intellisense

> used for generate code map → grammar compiler

[https://github.com/OmniSharp](https://github.com/OmniSharp)

Roslyn

[https://github.com/dotnet/roslyn](https://github.com/dotnet/roslyn)

## Important Feature

[https://docs.microsoft.com/en-us/dotnet/csharp/tour-of-csharp/features](https://docs.microsoft.com/en-us/dotnet/csharp/tour-of-csharp/features)

# Runtime

1. **System.Action**

![](Untitled-67848965-4466-4ea1-8a00-36fed38b2830.png)

use `(context, config)=>{METHOD}`

![](Untitled-d47b087f-0577-4ef8-8a90-0d40d3070bd4.png)

![](Untitled-aa7824c4-9536-4954-a95c-87ddd594cf08.png)

delegate相当于声明一个固定的函数类型的指针

1. **Attributes & Reflection**
- 为什么要用这种机制 有什么用？设计实例？
- 比如 `[Serializable]` 除了告知这个class可以serialize之外，它还有什么用吗？
- [https://docs.microsoft.com/zh-cn/dotnet/api/system.reflection.assembly?view=net-5.0](https://docs.microsoft.com/zh-cn/dotnet/api/system.reflection.assembly?view=net-5.0)

1. **UnitTest using MSTest**

[https://docs.microsoft.com/en-us/visualstudio/test/using-microsoft-visualstudio-testtools-unittesting-members-in-unit-tests?view=vs-2019](https://docs.microsoft.com/en-us/visualstudio/test/using-microsoft-visualstudio-testtools-unittesting-members-in-unit-tests?view=vs-2019)

1. **System.Linq.Expression**

![](Untitled-ec21d28d-e79e-4ac8-86a2-a36171bcfb5a.png)
- what is `Expression`1.cs` ?
- what's the diffs btw `Expression`1.cs` and `Expression.cs` ?
- linq?
- lambda expression?

**Dependency Inversion**

[https://docs.microsoft.com/en-us/dotnet/architecture/modern-web-apps-azure/architectural-principles#dependency-inversion](https://docs.microsoft.com/en-us/dotnet/architecture/modern-web-apps-azure/architectural-principles#dependency-inversion)

Lazy Initializer

**Singleton Design in C#**

1. basic c# singleton impletation

[https://csharpindepth.com/articles/singleton](https://csharpindepth.com/articles/singleton)

1. diffs btw singleton and static?

[https://stackoverflow.com/questions/519520/difference-between-static-class-and-singleton-pattern](https://stackoverflow.com/questions/519520/difference-between-static-class-and-singleton-pattern)

1. heap and stack — memory allocation
- ⭐ [https://stackoverflow.com/questions/33562199/static-class-memory-allocation-where-it-is-stored-c-sharp](https://stackoverflow.com/questions/33562199/static-class-memory-allocation-where-it-is-stored-c-sharp)
- [https://stackify.com/java-heap-vs-stack/](https://stackify.com/java-heap-vs-stack/)
- [https://www.geeksforgeeks.org/stack-vs-heap-memory-allocation/](https://www.geeksforgeeks.org/stack-vs-heap-memory-allocation/)
-