# JUnit

> How JUnit work

```java
invokeBeforeClass(Calculator.class); // @BeforeClass
for (Method testMethod : testMethods) {
    TestClass test = new TestClass();
    test.setUp(); // @Before
    testMethod.invoke(test); // @Test
    test.tearDown(); // @After 
}
invokeAfterClass(Calculator.class); // @AfterClass
```



## Annotation

```java
@Test
@Test(expected = NumberFormatException.class) // if throw NumberFormatException == OK, else fail
@Test(timeout = 1000) //time out test ==> fail when time out

Fixture
-------
@Before // for case eg input = new FileInputStream()
@After // for case eg input.close
@BeforeClass ==> field and method must be using static ==> this field will affect ALL @Test // for case eg build database 
@AfterClass ==> field and method must be using static ==> this field will affect ALL @Test // for case eg delete database
```



## Assert (fail when condition not matched)

```java
assertEquals(100, y)
assertEquals(3.1416, x, 0.0001) // last param is delta: +- difference accepted
assertArrayEqual({1,2,3},x)
assertNull(x)
assertTrue(x > 0)
assertFalse(x < 0)
assertNotEquals(100, x)
assertNotNull(x)
```



## Parameterized

> put the data that needed to be compared into collection.
>
> assertEquals(100, parameterized data)

> 1. @Parameters with static method: data(), returning a Collection
> 2. class with @RunWith(Parameterized.class)
> 3. Constructer parameter shall e same with Collection detail

```java
@RunWith(Parameterized.class)
public class AbsTest {
    @Parameters
    public static Collection<?> data() {
        return Arrays.asList(new Object[][] {
            { 0,0 }, { 1,1 }, { -1,1 }
        });
    }
    
    int input;
    int expected;
    
    public AbsTest(int input, int expected) {
        this.input = input;
        this.expected = expected;
    }
    
    @Test
    public void testAbs() {
        int r = Math.abs(this.input);
        assertEquals(this.expected, r);
    }
    
}
```