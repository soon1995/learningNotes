# Spring Data JPA (Jakarta Persistence API)

This is a learning notes from [bilibili -- 图灵课堂](https://www.bilibili.com/video/BV13Y411x7n9?p=3&spm_id_from=pageDriver&vd_source=a788bdd4d7cdd9dfe02852346d523cb9)

Start: 17/08/2022

End: 18/08/2022



> ORM -- Object Relational Mapping



## JPA Annotation

> if using hibernage.cfg.xml

```xml
<!-- since no more hbm.xml change from -->
<mapping resource="com/atguigu/....hbm,xml"/>

to 

<!-- to class -->
<mapping class="com.atguigu.entity.customer"/>

```

**Annotations:**

> If a method with FetchType.LAZY, there must be a @Transactional([readOnly = true]) (Throw LazyInitializationException when without). Which closing session **after** the method ends.

```java
-----
Class
-----
@Entity    
@Table(name="t_customer") // table name in database
    
-----
PK
-----
@Id
@GeneratedValue(strategy=GenerationType.IDENTITY)
@Column(name="id") // column name in database
    
------------
Normal Field
------------
@Column(name="name")
@Transient // 临时, in case does not hope to map this class field to database.
    
---------------
Optimistic Lock
---------------
@Version //(javax.persistence)
    
    
    
    
--------------------------
One To Many | Many To One (One way)
--------------------------
One Side
--------
@OneToMany(cascade = CascadeType.ALL)
@JoinColumn(name="customer_id")
    
Many Side
---------
    -
    
--------------------------
One To Many | Many To One (Two ways)
--------------------------
One Side
--------
@OneToMany(cascade = CascadeType.ALL)
@JoinColumn(name="customer_id") // Set<Order>
// Default FetchType = LAZY
    
Many Side
---------
@ManyToOne(cascade = CascadeType.ALL)
    
    
    
    
    
-------------
Many to Many (One way)
-------------
Side A
------
@ManyToMany(cascade = CascadeType.PERSIST)
@JoinTable(
    name="t_user_role", 
    joinColumns= {@JoinColumn(name="user_id")}, 
    inverseJoinColumns={@JoinColumn(name="role_id")})
    
    
-------------
Many to Many
-------------    
Side A
------
@ManyToMany(cascade = CascadeType.PERSIST)
@JoinTable(name="t_user_role", joinColumns=@JoinColumns(name="user_id"), inverseJoinColumns=@JoinColumn(name="role_id"))

Side B
------
@ManyToMany(mappedBy="roles")
    
    
    
    
    
    
--------------------
One To One (One Way)
--------------------
@OneToOne(cascade = CascadeType.ALL) 
//Tips: fetch = FetchType.LAZY if prefer lazy load, 
// orphanRemoval=true (default false) -> enable when set customer.setAccount(null); repository.save(customer), the account will be deleted in database
// optional=false (default true) -> this field cannot be null (throw Exception)
@JoinColumn(name="fk column Name")

--------------------
One To One (Both Ways)
--------------------
Point A
-------
@OneToOne(mappedBy = "customer(opposite class field)", // give up the foreign key
              cascade = CascadeType.ALL)
@JoinColumn(name="account_id")

Point B
-------
@OneToOne
@JoinColumn(name="customer_id")
```



----------------------------

**1a. Example of One To Many **(One Way)

Customer.class

```java
@OneToMany(cascade = CascadeType.ALL)
@JoinColumn(name="customer_id")
private List<Message> messages;
```

Message.class

```java
...
...
```



**1b. Example of One To Many** (Both Ways)

Customer.class

```java
...
    @OneToMany(cascade = CascadeType.ALL)
	@JoinColumn(name="customer_id")
    private List<Message> message;
...
```

Message.class

```java
...
    @ManyToOne
    private Customer customer;
...
```



-----------------------------

**2a. Example of Many to Many** (One Way)

User.class

```java
...
    @ManyToMany(cascade = CascadeType.PERSIST)
    @JoinTable(
            name="t_user_role",
            joinColumns= {@JoinColumn(name="user_id")},
            inverseJoinColumns={@JoinColumn(name="role_id")})
    private Set<Role> roles = new HashSet<>();
...
```

Message.class

```java
...
...
```



**2b. Example of Many to Many **(Both Ways)

User.class

```java
...
    @ManyToMany(cascade = CascadeType.PERSIST)
    @JoinTable(
            name="t_user_role",
            joinColumns= {@JoinColumn(name="user_id")},
            inverseJoinColumns={@JoinColumn(name="role_id")})
    private Set<Role> roles = new HashSet<>();
...
```

Role.class

```java
...
    @ManyToMany(mappedBy="roles")
    private Set<Customer> customers = new HashSet<>();
...
```



------------------------------------



**3a. Example of One to One **(One Way)

Customer.class

```java
...
   @OneToOne(cascade = CascadeType.ALL)
   @JoinColumn(name="account_id")
   private Account account;
...
```

Account.class

```java
...
...
```



**3b. Example of One to One **(Both Ways)

Customer.class

```java
...
   @OneToOne(mappedBy = "customer", // give up the foreign key
    			cascade = CascadeType.ALL)
   @JoinColumn(name="account_id")
   private Account account;
...
```

Account.class

```java
...
   @OneToOne
   @JoinColumn(name="customer_id")
   private Customer customer;
...
```



> JPA Primary Key Strategy
>
> - IDENTIY: database auto_increment (for mysql)
> - SEQUENCE: database sequence (for oracle)
> - TABLE: JPA generated a table to keep latest id for tables
> - AUTO



## Original JPA

Dependency

```xml
<dependency>
    <groupId>junit</groupId>
    <artifactId>junit</artifactId>
    <version>4.13.1</version>
    <scope>test</scope>
</dependency>
<dependency>
    <groupId>org.hibernate</groupId>
    <artifactId>hibernate-entitymanager</artifactId>
    <version>5.6.10.Final</version>
</dependency>
        <dependency>
            <groupId>org.hibernate</groupId>
            <artifactId>hibernate-c3p0</artifactId>
            <version>5.6.10.Final</version>
        </dependency>
<dependency>
    <groupId>mysql</groupId>
    <artifactId>mysql-connector-java</artifactId>
    <version>8.0.29</version>
</dependency>
```

resources/META-INF/persistence.xml

```xml
<?xml version="1.0" encoding="UTF-8"?>
<persistence version="2.0" xmlns="http://java.sun.com/xml/ns/persistence">
    <!-- this name will tally with EntityManagerFactory -->
    <persistence-unit name="hibernateJPA" transaction-type="RESOURCE_LOCAL">
        <provider>org.hibernate.jpa.HibernatePersistenceProvider</provider>

        <class>com.example.entity.Customer1</class>

        <properties>
            <property name="javax.persistence.jdbc.driver" value="com.mysql.cj.jdbc.Driver"/>
            <property name="javax.persistence.jdbc.url" value="jdbc:mysql://localhost:3306/test_hibernate"/>
            <property name="javax.persistence.jdbc.user" value="root"/>
            <property name="javax.persistence.jdbc.password" value="alsoon"/>

            <property name="hibernate.show_sql" value="true"/>
            <property name="hibernate.format_sql" value="true"/>
            <property name="hibernate.dialect" value="org.hibernate.dialect.MySQL5Dialect"/>
            <property name="hibernate.hbm2ddl.auto" value="update"/>

            <property name="hibernate.c3p0.max_size" value="10"/>
            <property name="hibernate.c3p0.min_size" value="5"/>
            <property name="c3p0.acquire_increment" value="2"/>
            <property name="c3p0.timeout" value="2000"/>
            <property name="c3p0.idle_test_period" value="2000"/>
            <property name="c3p0.max_statements" value="10"/>
        </properties>

    </persistence-unit>
</persistence>
```

Example

```java
import com.example.entity.Customer1;
import com.example.temp.Customer;
import org.hibernate.Session;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;

import javax.persistence.EntityManager;
import javax.persistence.EntityManagerFactory;
import javax.persistence.EntityTransaction;
import javax.persistence.Persistence;

public class TestJPA {
    EntityManagerFactory emf;
    EntityManager em;
    EntityTransaction tx;

    @Before
    public void init() {
        emf = Persistence.createEntityManagerFactory("hibernateJPA");
        em = emf.createEntityManager();
        tx = em.getTransaction();
    }

    @After
    public void destroy() {
        em.close();
        emf.close();
    }

    @Test
    public void testSave() {
        tx.begin();
        try {
            Customer1 customer = new Customer1();
            customer.setCustomerName("张三");

            em.persist(customer);
            tx.commit();
        } catch (Exception e) {
            tx.rollback();
        }
    }

    @Test
    public void testFind() {
        tx.begin();
        try {
            Customer1 customer = em.find(Customer1.class, 1);
            System.out.println("=========="); // sql execute before this (no lazy)
            System.out.println(customer.getCustomerName());
            tx.commit();
        } catch (Exception e) {
            tx.rollback();
        }
    }

    @Test
    public void testLazyFind() {
        tx.begin();
        try {
            Customer1 customer = em.getReference(Customer1.class, 1);
            System.out.println("=========="); // sql execute AFTER this (lazy)
            System.out.println(customer.getCustomerName());
            tx.commit();
        } catch (Exception e) {
            tx.rollback();
        }
    }


    @Test
    public void testSaveOrUpdate() {
        tx.begin();
        try {
            Customer1 customer1 = new Customer1();
            customer1.setCustomerId(1);
            customer1.setCustomerName("李四");
            // select sql x 1, update sql x 1
            em.merge(customer1);
            tx.commit();
        } catch (Exception e) {
            tx.rollback();
        }
    }

    @Test
    public void testSaveOrUpdate2() {
        tx.begin();
        try {
            Customer1 customer1 = new Customer1();
            customer1.setCustomerName("李四");
            // select insert sql x 1
            em.merge(customer1);
            tx.commit();
        } catch (Exception e) {
            tx.rollback();
        }
    }

    @Test
    public void testJPQL() {
        tx.begin();
        try {
            em.createQuery("UPDATE Customer1 set customerName=:name where customerId=:id")
                    .setParameter("name", "王五")
                    .setParameter("id",1)
                    .executeUpdate();
            tx.commit();
        } catch (Exception e) {
            tx.rollback();
        }
    }

    @Test
    public void testSQL() {
        tx.begin();
        try {
            em.createNativeQuery("UPDATE t_customer set name=:name where id=:id")
                    .setParameter("name", "张三")
                    .setParameter("id",1)
                    .executeUpdate();
            tx.commit();
        } catch (Exception e) {
            tx.rollback();
        }
    }

    @Test
    public void testRemove() {
        tx.begin();
        try {
            // must get first instead of new customer, because cannot remove detached status object
            Customer1 customer1 = em.find(Customer1.class, 1);
            em.remove(customer1);
            tx.commit();
        } catch (Exception e) {
            tx.rollback();
        }
    }
}
```



## Setting Up Spring-Data-jpa with Spring

### DependencyManagement

```xml
<dependencyManagement>
	<dependencies>
        <!-- help to manage the version of the spring data eg JPA, redis ... -->
    	<dependency>
        	<groupId>org.springframework.data</groupId>
            <artifactId>spring-data-bom</artifactId>
			<version>2021.1.0</version>
            <scope>import</scope>
            <type>pom</type>
        </dependency>
    </dependencies>
</dependencyManagement>
```

### Dependency in module

```xml
<dependency>
	<groupId>org.springframework.data</groupID>
    <artifactId>spring-data-jpa</artifactId>
</dependency>
<!-- 
includes:
spring-data-commons
spring-orm
spring-context
spring-aop
spring-tx
spring-beans
spring-core
slf4j
-->
<dependency>
    <groupId>junit</groupId>
    <artifactId>junit</artifactId>
    <version>4.13.1</version>
    <scope>test</scope>
</dependency>
<dependency>
    <groupId>org.hibernate</groupId>
    <artifactId>hibernate-entitymanager</artifactId>
    <version>5.6.10.Final</version>
</dependency>
<dependency>
    <groupId>org.hibernate</groupId>
    <artifactId>hibernate-c3p0</artifactId>
    <version>5.6.10.Final</version>
</dependency>
<dependency>
    <groupId>org.springframework</groupId>
    <artifactId>spring-test</artifactId>
    <version>5.3.10</version>
    <scope>test</scope>
</dependency>
<dependency>
    <groupId>mysql</groupId>
    <artifactId>mysql-connector-java</artifactId>
    <version>8.0.29</version>
</dependency>
```

### Configuration

- xml please refer [here](https://www.bilibili.com/video/BV13Y411x7n9?p=11)

- java config

```java
package com.alsoon.config;

import com.mchange.v2.c3p0.ComboPooledDataSource;
import com.mchange.v2.c3p0.jboss.C3P0PooledDataSource;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.jpa.repository.config.EnableJpaRepositories;
import org.springframework.orm.jpa.JpaTransactionManager;
import org.springframework.orm.jpa.LocalContainerEntityManagerFactoryBean;
import org.springframework.orm.jpa.vendor.HibernateJpaVendorAdapter;
import org.springframework.transaction.PlatformTransactionManager;
import org.springframework.transaction.annotation.EnableTransactionManagement;

import javax.naming.NamingException;
import javax.persistence.EntityManagerFactory;
import javax.sql.DataSource;
import java.beans.PropertyVetoException;

@Configuration
@EnableJpaRepositories(basePackages="com.alsoon.repository")
@EnableTransactionManagement
@EnableJpaAuditing
public class SpringDataJPAConfig {
    @Bean
    public DataSource dataSource() {
        ComboPooledDataSource dataSource = new ComboPooledDataSource();
        try {
            dataSource.setUser("xxxxxxxxxxx");
            dataSource.setPassword("xxxxxxxxx");
            dataSource.setDriverClass("com.mysql.cj.jdbc.Driver");
            dataSource.setJdbcUrl("jdbc:mysql://localhost:3306/test_hibernate?serverTimezone=GMT%2B8");
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
        return dataSource;
    }

    @Bean
    public LocalContainerEntityManagerFactoryBean entityManagerFactory() {
        HibernateJpaVendorAdapter vendorAdapter = new HibernateJpaVendorAdapter();
        vendorAdapter.setGenerateDdl(true);
        vendorAdapter.setShowSql(true);

        LocalContainerEntityManagerFactoryBean factory = new LocalContainerEntityManagerFactoryBean();
        factory.setJpaVendorAdapter(vendorAdapter);
        factory.setPackagesToScan("com.alsoon.pojo");
        factory.setDataSource(dataSource());

        return factory;
    }

    @Bean
    public PlatformTransactionManager transactionManager(EntityManagerFactory entityManagerFactory) {
        JpaTransactionManager txManager = new JpaTransactionManager();
        txManager.setEntityManagerFactory(entityManagerFactory);
        return txManager;
    }
    
    /**
    * This is the createBy field in database
    * <Type> is refer to the Pojo field, might be String, might be User ....
    */
    @Bean
    public AuditorAware<String> auditorAware() {
        return () -> Optional.of("testing123");
    }
}

```



## Repository

- CrudRepository: Extends Repository
- PagingAndSortingRepository: Extends CrudRepository
- JpaRepository: Extends PagingAndSortingRepository
- Customize --- extends xxxxRepository

```java
package com.alsoon.repository;


import com.alsoon.pojo.Customer;
import org.springframework.data.repository.CrudRepository;

public interface CustomerRepository extends CrudRepository<Customer, Long> {
}

```

### CrudRepository

Refer [Spring-data Doc](https://docs.spring.io/spring-data/commons/docs/current/api/org/springframework/data/repository/CrudRepository.html)

| Modifier and Type          | Method                                      | Description                                                  |
| -------------------------- | ------------------------------------------- | ------------------------------------------------------------ |
| `long`                     | `count()`                                   | Returns the number of entities available.                    |
| `void`                     | `delete(T entity)`                          | Deletes a given entity.                                      |
| `void`                     | `deleteAll()`                               | Deletes all entities managed by the repository.              |
| `void`                     | `deleteAll(Iterable<? extends T> entities)` | Deletes the given entities.                                  |
| `void`                     | `deleteAllById(Iterable<? extends ID> ids)` | Deletes all instances of the type `T` with the given IDs     |
| `void`                     | `deleteById(ID id)`                         | Deletes the entity with the given id.                        |
| `boolean`                  | `existsById(ID id)`                         | Returns whether an entity with the given id exists.          |
| `Iterable<T>`              | `findAll()`                                 | Returns all instances of the type.                           |
| `Iterable<T>`              | `findAllById(Iterable<ID> ids)`             | Returns all instances of the type `T` with the given IDs.    |
| `Optional<T>`              | `findById(ID id)`                           | Retrieves an entity by its id.                               |
| `<S extends T>S`           | `save(S entity)`                            | Saves a given entity.                                        |
| `<S extends T>Iterable<S>` | `saveAll(Iterable<S> entities)`             | Saves all given entities. Be noted this is not batch insert, object are inserted one by one |

```java
import com.alsoon.config.SpringDataJPAConfig;
import com.alsoon.pojo.Customer;
import com.alsoon.repository.CustomerRepository;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.test.context.ContextConfiguration;
import org.springframework.test.context.junit4.SpringJUnit4ClassRunner;

import java.util.Optional;

@ContextConfiguration(classes = SpringDataJPAConfig.class)
@RunWith(SpringJUnit4ClassRunner.class)
public class Test {
    @Autowired
    CustomerRepository repository;

    @org.junit.Test
    public void testQuery() {
        // Optional -> jdk8 new characteristic to prevent Null Pointer
        Optional<Customer> byId = repository.findById(1L);
        System.out.println(byId.get());
    }

    @org.junit.Test
    public void testInsert() {
        Customer customer = new Customer();
        customer.setCustName("李四");

        repository.save(customer);
    }

    @org.junit.Test
    public void testUpdate() {
        Customer customer = new Customer();
        customer.setCustId(4L);
        customer.setCustName("张三");

        repository.save(customer);
    }

    @org.junit.Test
    public void testDelete() {
        Customer customer = new Customer();
        customer.setCustId(4L);

        repository.delete(customer);
    }
}
```



### PagingAndSortingRepository

| Modifier and Type | Method                       | Description                                                  |
| ----------------- | ---------------------------- | ------------------------------------------------------------ |
| `Page<T>`         | `findAll(Pageable pageable)` | Returns a [`Page`](https://docs.spring.io/spring-data/commons/docs/current/api/org/springframework/data/domain/Page.html) of entities meeting the paging restriction provided in the [`Pageable`](https://docs.spring.io/spring-data/commons/docs/current/api/org/springframework/data/domain/Pageable.html) object. |
| `Iterable<T>`     | `findAll(Sort sort)`         | Returns all entities sorted by the given options.            |

```java
@ContextConfiguration(classes = SpringDataJPAConfig.class)
@RunWith(SpringJUnit4ClassRunner.class)
public class PagingRepositoryTest {
    @Autowired
    CustomerRepository repository;

    @org.junit.Test
    public void testPage() {
        // PageRequest of(int page, int size) page start from 0, must not be negative
        // PageRequest.of(int page, int size, Sort) can be use if want order by
        Page<Customer> datas = repository.findAll(PageRequest.of(0, 1));
        List<Customer> content = datas.getContent();
        int totalPages = datas.getTotalPages();
        long totalElements = datas.getTotalElements();

        System.out.println(content);
        System.out.println(totalPages);
        System.out.println(totalElements);
    }
    
    @org.junit.Test
    public void testSort() {
        Sort.TypedSort<Customer> customer = Sort.sort(Customer.class);
        Sort sort = customer.by(Customer::getCustId).descending().and(customer.by(Customer::getCustName).ascending());
        Iterable<Customer> all = repository.findAll(sort);
        System.out.println(all);
    }
}

```





## Customize Operation

- JPQL
- Method's Name
- Query By Example (Achieve Dynamic Query)
- Specifications (Achieve Dynamic Query)
- QueryDSL (Achieve Dynamic Query)



### JPQL

- @Query("JPQL")
- Install plugin (JPABuddy [IDEA]) ==> have tips when typing JPQL
- Native SQL: @Query(value="Native SQL", native=true) if complicated SQL to be executed

```java
public interface CustomerRepository extends PagingAndSortingRepository<Customer, Long> {

    @Query("FROM Customer WHERE custName=?1")
    Customer findCustomerByCustName(String custName);


    @Query("FROM Customer WHERE custName=:custName")
    List<Customer> findCustomersByCustName(@Param("custName") String custName);
    
    @Transactional
    @Modifying
    @Query("UPDATE Customer c SET c.custName=:custName WHERE c.custId=:id")
    int updateCustomerById(@Param("custName") String custName, @Param("id") Long id);
    
    @Transactional
    @Modifying
    @Query("DELETE Customer c WHERE c.custId=:id")
    int deleteCustomerById(@Param("id") Long id);
    
    
    // JPQL Not support Insert, only support insert .. select
    @Transactional
    @Modifying
    @Query("INSERT INTO Customer(custName) SELECT c.custName FROM Customer c WHERE c.custId=:id")
    int insertCustomerBySelect(@Param("id") Long id);
    
    @Transactional
    @Modifying
    @Query(nativeQuery = true, value="INSERT INTO t_customer (`name`) VALUES(:name)")
    int insertCustomerByNative(@Param("name") String name);

}
```

```java
@ContextConfiguration(classes = SpringDataJPAConfig.class)
@RunWith(SpringJUnit4ClassRunner.class)
public class JPQLTest {
    @Autowired
    CustomerRepository repository;

    @org.junit.Test
    public void testFind() {
        Customer customer = repository.findCustomerByCustName("王五");
        System.out.println(customer);

        List<Customer> customer2 = repository.findCustomersByCustName("李四");
        System.out.println(customer2);
    }
    
    @org.junit.Test
    public void testUpdate() {
        // throw error if no @Transaction, normally this annotation appeared in Services instead of Repository
        // throw error if no @Modifying, inform this is a Insert, Update, Delete sql to SpringDataJPA
        System.out.println(repository.updateCustomerById("王五", 2L));
    }
    
    @org.junit.Test
    public void testDelete() {
        // ditto
        System.out.println(repository.deleteCustomerById(3L));
    }
    
    @org.junit.Test
    public void testInsertBySelect() {
        // ditto
        System.out.println(repository.insertCustomerBySelect(2L));
    }
    
    @org.junit.Test
    public void testInsert() {
        // ditto
        System.out.println(repository.insertCustomerByNative("Ahlong"));
    }
}
```



### Method's Name

> Only support query (including count) & delete. Do not support update and insert.

refer [Spring](https://docs.spring.io/spring-data/jpa/docs/current/reference/html/#repository-query-keywords) for more detail information

#### Example

```java
public interface CustomMethodRepository extends PagingAndSortingRepository<Customer, Long> {

    List<Customer> findByCustName(String name);

    boolean existsCustomerByCustId(Long id);

    //Be noted this one cost 1 select, and multiple delete sql
    @Transactional
    @Modifying
    int deleteAllByCustName(String name);

    List<Customer> findByCustNameLike(String pattern);

}

```

```java
@ContextConfiguration(classes = SpringDataJPAConfig.class)
@RunWith(SpringJUnit4ClassRunner.class)
public class MethodNameTest {
    @Autowired
    CustomMethodRepository repository;

    @org.junit.Test
    public void testFind() {
        System.out.println(repository.findByCustName("李四"));
    }

    @org.junit.Test
    public void testExist() {
        System.out.println(repository.existsCustomerByCustId(5L));
    }

    @org.junit.Test
    public void testDelete() {
        System.out.println(repository.deleteAllByCustName("李四"));
    }

    @org.junit.Test
    public void testLike() {
        System.out.println(repository.findByCustNameLike("%王%"));
    }
}
```



#### Query Subject Keywords

| Keyword                                                      | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `find…By`, `read…By`, `get…By`, `query…By`, `search…By`, `stream…By` | General query method returning typically the repository type, a `Collection` or `Streamable` subtype or a result wrapper such as `Page`, `GeoResults` or any other store-specific result wrapper. Can be used as `findBy…`, `findMyDomainTypeBy…` or in combination with additional keywords. |
| `exists…By`                                                  | Exists projection, returning typically a `boolean` result.   |
| `count…By`                                                   | Count projection returning a numeric result.                 |
| `delete…By`, `remove…By`                                     | Delete query method returning either no result (`void`) or the delete count. |
| `…First<number>…`, `…Top<number>…`                           | Limit the query results to the first `<number>` of results. This keyword can occur in any place of the subject **between** `find` (and the other keywords) and `by`, eg findFirst2By... \| findFirstBy == findFirst1By |
| `…Distinct…`                                                 | Use a distinct query to return only unique results. Consult the  store-specific documentation whether that feature is supported. This  keyword can occur in any place of the subject between `find` (and the other keywords) and `by`. |



#### Query Predicate Keywords

| Keyword             | Sample                                                    | JPQL snippet                                                 |
| ------------------- | --------------------------------------------------------- | ------------------------------------------------------------ |
| `And`               | `findByLastnameAndFirstname`                              | `… where x.lastname = ?1 and x.firstname = ?2`               |
| `Or`                | `findByLastnameOrFirstname`                               | `… where x.lastname = ?1 or x.firstname = ?2`                |
| `Is,Equals`         | `findByFirstname,findByFirstnameIs,findByFirstnameEquals` | `… where x.firstname = 1?`                                   |
| `Between`           | `findByStartDateBetween`                                  | `… where x.startDate between 1? and ?2`                      |
| `LessThan`          | `findByAgeLessThan`                                       | `… where x.age < ?1`                                         |
| `LessThanEqual`     | `findByAgeLessThanEqual`                                  | `… where x.age <= ?1`                                        |
| `GreaterThan`       | `findByAgeGreaterThan`                                    | `… where x.age > ?1`                                         |
| `GreaterThanEqual`  | `findByAgeGreaterThanEqual`                               | `… where x.age >= ?1`                                        |
| `After`             | `findByStartDateAfter`                                    | `… where x.startDate > ?1`                                   |
| `Before`            | `findByStartDateBefore`                                   | `… where x.startDate < ?1`                                   |
| `IsNull`            | `findByAgeIsNull`                                         | `… where x.age is null`                                      |
| `IsNotNull,NotNull` | `findByAge(Is)NotNull`                                    | `… where x.age not null`                                     |
| `Like`              | `findByFirstnameLike`                                     | `… where x.firstname like ?1` **THE PARAM SHOULD WITH %**    |
| `NotLike`           | `findByFirstnameNotLike`                                  | `… where x.firstname not like ?1`                            |
| `StartingWith`      | `findByFirstnameStartingWith`                             | `… where x.firstname like ?1` (parameter bound with appended `%`) |
| `EndingWith`        | `findByFirstnameEndingWith`                               | `… where x.firstname like ?1` (parameter bound with prepended `%`) |
| `Containing`        | `findByFirstnameContaining`                               | `… where x.firstname like ?1` (parameter bound wrapped in `%`) |
| `OrderBy`           | `findByAgeOrderByLastnameDesc`                            | `… where x.age = ?1 order by x.lastname desc`                |
| `Not`               | `findByLastnameNot`                                       | `… where x.lastname <> ?1`                                   |
| `In`                | `findByAgeIn(Collection<Age> ages)`                       | `… where x.age in ?1`                                        |
| `NotIn`             | `findByAgeNotIn(Collection<Age> age)`                     | `… where x.age not in ?1`                                    |
| `True`              | `findByActiveTrue()`                                      | `… where x.active = true`                                    |
| `False`             | `findByActiveFalse()`                                     | `… where x.active = false`                                   |
| `IgnoreCase`        | `findByFirstnameIgnoreCase`                               | `… where UPPER(x.firstame) = UPPER(?1)`                      |



#### Query predicate modifier keywords

| Keyword                            | Description                                                  |
| ---------------------------------- | ------------------------------------------------------------ |
| `IgnoreCase`, `IgnoringCase`       | Used with a predicate keyword for case-insensitive comparison. |
| `AllIgnoreCase`, `AllIgnoringCase` | Ignore case for all suitable properties. Used somewhere in the query method predicate. |
| `OrderBy…`                         | Specify a static sorting order followed by the property path and direction (e. g. `OrderByFirstnameAscLastnameDesc`). |



### Query By Example

- only support query
  - only support String : start/contains/ends/regex 
  - or other type for exact match eg. failing to do date range query
  - not working with embedded condition eg [firstname=?0 or **(firstname =?1 and lastname=?2)**].

```java
public interface CustomerQBERepository extends PagingAndSortingRepository<Customer, Long>, QueryByExampleExecutor<CustomerQBERepository> {

}

```

```java
@ContextConfiguration(classes = SpringDataJPAConfig.class)
@RunWith(SpringJUnit4ClassRunner.class)
public class QBETest {

    @Autowired
    CustomerQBERepository repository;

    /**
     * example: customerName customerAddr dynamic query
     */
    @Test
    public void test01() {
        Customer customer = new Customer();
        customer.setCustName("王五");
        customer.setCustAddr("beijing");

        // The example is from springframework package instead of Hibernate
        Example example = Example.of(customer);

        List<Customer> customerList = (List<Customer>) repository.findAll(example);

        System.out.println(customerList);
    }

    /**
     * example: customerName customerAddr dynamic query -- ignore the condition of custAddr
     */
    @Test
    public void test02() {
        Customer customer = new Customer();
        customer.setCustName("王五");
        customer.setCustAddr("beijing");

        ExampleMatcher matcher = ExampleMatcher.matching()
                .withIgnorePaths("custAddr");

        // The example is from springframework package instead of Hibernate
        Example example = Example.of(customer, matcher);

        List<Customer> customerList = (List<Customer>) repository.findAll(example);

        System.out.println(customerList);
    }

    /**
     * example: customerName customerAddr dynamic query -- case insensitive
     */
    @Test
    public void test03() {
        Customer customer = new Customer();
        customer.setCustName("王五");
        customer.setCustAddr("bei");

        ExampleMatcher matcher = ExampleMatcher.matching()
                .withIgnoreCase()
                .withMatcher("custAddr", matcher1 -> matcher1.startsWith().ignoreCase())
                ;
        //.withIgnoreCase("custAddr") --> only ignorecase for this field
        //.withStringMatcher(ExampleMatcher.StringMatcher.STARTING) --> to all field

        // The example is from springframework package instead of Hibernate
        Example example = Example.of(customer, matcher);

        List<Customer> customerList = (List<Customer>) repository.findAll(example);

        System.out.println(customerList);
    }

    @Test
    public void SpringExample() {
        Customer customer = new Customer();
        customer.setCustName("Dave");

        // withIncludeNullValues means where custName = "Dave and (custId is null) and (custAddr is null)
        ExampleMatcher matcher = ExampleMatcher.matching()
                .withIgnorePaths("custName")
                .withIncludeNullValues()
                .withStringMatcher(ExampleMatcher.StringMatcher.ENDING);

        Example<Customer> example = Example.of(customer, matcher);
    }

}

```



#### ExampleMatcher

| Setting              | Scope                              |
| -------------------- | ---------------------------------- |
| Null-handling        | `ExampleMatcher`                   |
| String matching      | `ExampleMatcher` and property path |
| Ignoring properties  | Property path                      |
| Case sensitivity     | `ExampleMatcher` and property path |
| Value transformation | Property path                      |

#### StringMatcher

| Matching                        | Logical result                              |
| ------------------------------- | ------------------------------------------- |
| `DEFAULT` (case-sensitive)      | `firstname = ?0`                            |
| `DEFAULT` (case-insensitive)    | `LOWER(firstname) = LOWER(?0)`              |
| `EXACT` (case-sensitive)        | `firstname = ?0`                            |
| `EXACT` (case-insensitive)      | `LOWER(firstname) = LOWER(?0)`              |
| `STARTING` (case-sensitive)     | `firstname like ?0 + '%'`                   |
| `STARTING` (case-insensitive)   | `LOWER(firstname) like LOWER(?0) + '%'`     |
| `ENDING` (case-sensitive)       | `firstname like '%' + ?0`                   |
| `ENDING` (case-insensitive)     | `LOWER(firstname) like '%' + LOWER(?0)`     |
| `CONTAINING` (case-sensitive)   | `firstname like '%' + ?0 + '%'`             |
| `CONTAINING` (case-insensitive) | `LOWER(firstname) like '%' + LOWER(?0) + '` |



### Specification

> Not support grouping

```java
public interface CustomSpecificationsRepository extends PagingAndSortingRepository<Customer, Long>, JpaSpecificationExecutor<Customer> {
}
```

```java
@ContextConfiguration(classes = SpringDataJPAConfig.class)
@RunWith(SpringJUnit4ClassRunner.class)
public class SpecificationsTest {
    @Autowired
    private CustomSpecificationsRepository repository;

    @Test
    public void test() {

        Customer customer = new Customer();
        customer.setCustName("王五");
        customer.setCustId(1L);
//        customer.setCustAddr("asd");


        List<Customer> result = repository.findAll((Specification<Customer>) (root, query, criteriaBuilder) -> {
            // root : from Customer, get Column
            // criteriaBuilder : where ... condition (>< in..)
            // query : orderBy, where

            /*
            find customer range
            id >
            addr =
            name in
             */

            // 1. get field from root
            Path<Long> custId = root.get("custId");
            Path<String> custName = root.get("custName");
            Path<String> custAddr = root.get("custAddr");

            //2. set conditions by criteriaBuilder
            List<Predicate> list = new ArrayList<>();

            if (customer.getCustId() != null && customer.getCustId() > -1) {
                list.add(criteriaBuilder.ge(custId, 2L));
            }

            if (!StringUtils.isEmpty(customer.getCustName())) {
                //Predicate predicate3 = criteriaBuilder.like(custName, "%王%");
                CriteriaBuilder.In<String> in = criteriaBuilder.in(custName);
                in.value("王五").value("张三");
                list.add(in);

            }

            if(!StringUtils.isEmpty(customer.getCustAddr())) {
                list.add(criteriaBuilder.equal(custAddr, "BEIJING"));
            }

            Predicate and = criteriaBuilder.and(list.toArray(new Predicate[list.size()]));

            // 3. (Optional) orderBy 组合条件
            Order order = criteriaBuilder.desc(custId);

            return query.where(and).orderBy(order).getRestriction();
        });

        System.out.println(result);
    }
}
```



### QueryDsl (Third Party Framework)

Dependency

```xml
<dependency>
    <groupId>com.querydsl</groupId>
    <artifactId>querydsl-jpa</artifactId>
    <version>5.0.0</version>
</dependency>
```

Plugin to generate "Q" class

```xml
    <build>
        <plugins>
            <plugin>
                <groupId>com.mysema.maven</groupId>
                <artifactId>apt-maven-plugin</artifactId>
                <version>1.1.3</version>
                <dependencies>
                    <dependency>
                        <groupId>com.querydsl</groupId>
                        <artifactId>querydsl-apt</artifactId>
                        <version>${querydsl.version}</version>
                    </dependency>
                </dependencies>
                <executions>
                    <execution>
                        <phase>generate-sources</phase>
                        <goals>
                            <goal>process</goal>
                        </goals>
                        <configuration>
                            <outputDirectory>target/generated-sources/queries</outputDirectory>
                            <processor>com.querydsl.apt.jpa.JPAAnnotationProcessor</processor>
                            <options>
                                <querydsl.entityAccessors>true</querydsl.entityAccessors>
                                <querydsl.useFields>false</querydsl.useFields>
                            </options>
                        </configuration>
                    </execution>
                </executions>
            </plugin>
        </plugins>
    </build>
```

Tips: to reflect immediately -> maven -> compile

Tips: in project structure set target/generated-sources/queries as sources (instead of excluded) to use 'Q' class 

```java
public interface CustomQueryDSLRepository extends PagingAndSortingRepository<Customer, Long>, QuerydslPredicateExecutor<Customer> {
}

```

```java
@ContextConfiguration(classes = SpringDataJPAConfig.class)
@RunWith(SpringJUnit4ClassRunner.class)
public class QueryDSLTest {

    @Autowired
    CustomQueryDSLRepository repository;
    
     // 解决线程安全问题，类似于Autowired！@Bean默认是单例的，JPA就有线程安全问题！
     @PersistenceContext
     EntityManager entityManager;

    @Test
    public void test01() {
        QCustomer customer = QCustomer.customer;

        // query by id
        BooleanExpression eq = customer.custId.eq(2L);

        System.out.println(repository.findOne(eq));
    }

    @Test
    public void test02() {
        QCustomer customer = QCustomer.customer;
        BooleanExpression condition = customer.custName.in("王五", "李四")
                .and(customer.custId.goe(2L))
                .and(customer.custAddr.likeIgnoreCase("%bei%"));

        System.out.println(repository.findAll(condition));
    }
    
    @Test
    public void dynamicTest() {
        Customer customer = new Customer();
        customer.setCustName("王五");
        customer.setCustAddr("abc");
        customer.setCustId(1L);


        QCustomer qCustomer = QCustomer.customer;
        // 1=1 condition to BooleanExpression always work
        BooleanExpression ex = qCustomer.isNotNull().or(qCustomer.isNull());

        if (customer.getCustId() != null && customer.getCustId() > -1) {
            ex = ex.and(qCustomer.custId.goe(2L));
        }

        if (!StringUtils.isEmpty(customer.getCustName())) {
            ex = ex.and(qCustomer.custName.in("王五", "李四"));
        }

        if(!StringUtils.isEmpty(customer.getCustAddr())) {
            ex = ex.and(qCustomer.custAddr.likeIgnoreCase("%bei%"));
        }

        System.out.println(repository.findAll(ex));
    }
    
    /**
      * 自定义列查询、分组
      * 需要使用原生态的方式！(Specification)
      */
     @Test
     public void test_customize(){
         JPAQueryFactory factory = new JPAQueryFactory(entityManager);

         QCustomer customer = QCustomer.customer;

         // select id,custName from ...
         // 构建基于QueryDsl的查询
         JPAQuery<Tuple> tupleJPAQuery = factory.select(customer.id, customer.custName)
                 .from(customer)
                 .where(customer.id.eq(1L))
                 .orderBy(customer.id.desc());

         // 执行查询
         List<Tuple> fetch = tupleJPAQuery.fetch();

         // 处理返回数据
         for (Tuple tuple : fetch) {
             System.out.println(tuple.get(customer.id));
             System.out.println(tuple.get(customer.custName));
         }
     }
    
    /**
      * 自定义列查询、分组
      * 需要使用原生态的方式！(Specification)
      * 通过Repository进行查询，列、表都是固定的！
      */
     @Test
     public void test_customize_list(){
         JPAQueryFactory factory = new JPAQueryFactory(entityManager);
 
         QCustomer customer = QCustomer.customer;
 
         // select id,custName from ...
         // 构建基于QueryDsl的查询
         JPAQuery<Long> longJPAQuery = factory.select(customer.id.sum())
                 .from(customer)
                 //.where(customer.id.eq(1L))
                 .orderBy(customer.id.desc());
 
         // 执行查询
         List<Long> fetch = longJPAQuery.fetch();
 
         // 处理返回数据
         for (Long tuple : fetch) {
             System.out.println(tuple);
         }
     }
}
```



## Audit Aware

> In case for some table we need to have 
>
> 1. createdBy
> 2. lastModifiedBy
> 3. createdDate
> 4. lastModifiedDate



1. **Dependency**

   ```xml
   <dependency>
       <groupId>org.springframework</groupId>
       <artifactId>spring-aspects</artifactId>
       <version>5.3.13</version>
   </dependency>
   ```

   

2. **Config**

   ```java
    import org.springframework.data.domain.AuditorAware;
    import java.util.Optional;
    
    /**
     * 监听
     * @CreatedBy
     * @LastModifiedBy
     * 自动注入用户名
     */
   @Configuration
    public class AuditorAwareConfig implements AuditorAware<String> {
        @Override
        public Optional<String> getCurrentAuditor() {
            return Optional.empty();
        }
    }
    
    ----------------------------------------------
        OR
    ----------------------------------------------
    
   @Configuration
   @EnableJpaRepositories(basePackages="com.alsoon.repository")
   @EnableTransactionManagement
   @EnableJpaAuditing //**************
   public class SpringDataJPAConfig {   
       ....
    
       /**
         * 这是JavaConfig方式！
         * AuditorAware 返回当前用户
         * @return
         */
           @Bean
           public AuditorAware<String> auditorAware() {
           return () -> Optional.of("testing");
       }
   }
   ```

   

2. **Pojo Class**

   ```java
   @EntityListeners(AuditingEntityListener.class)//***********
   @Entity
   @Table(...)
   @Data
   public class ImPojo {
       ...
           
        @CreatedBy
        String createdBy;
    
        @LastModifiedBy
        String modifiedBy;
    
        @Temporal(TemporalType.TIMESTAMP)
        @CreatedDate
        protected Date dateCreated = new Date();
    
        @Temporal(TemporalType.TIMESTAMP)
        @LastModifiedDate
        protected Date dateModified = new Date();
   }
   ```



## Setting Up Spring-Data-jpa with Spring Boot

1. **Dependency**

   ```xml
   <dependencies>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-data-jpa</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>mysql</groupId>
            <artifactId>mysql-connector-java</artifactId>
            <scope>runtime</scope>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
            <exclusions>
                <exclusion>
                    <groupId>org.junit.vintage</groupId>
                    <artifactId>junit-vintage-engine</artifactId>
                </exclusion>
            </exclusions>
        </dependency>
    </dependencies>
   ```

2. **application.properties**

   > Other properties can refer the [video 12:20 (In Chinese)](https://www.bilibili.com/video/BV13Y411x7n9?p=38&vd_source=a788bdd4d7cdd9dfe02852346d523cb9)

   ```properties
    # 应用名称
    spring.application.name=04-springdata-jpa-springboot
    # 数据库驱动：
    spring.datasource.driver-class-name=com.mysql.cj.jdbc.Driver
    # 数据源名称
    spring.datasource.name=defaultDataSource
    # 数据库连接地址
    spring.datasource.url=jdbc:mysql://localhost:3306/spring_data?serverTimezone=UTC
    # 数据库用户名&密码：
    spring.datasource.username=root
    spring.datasource.password=123456
    # 应用服务 WEB 访问端口
    server.port=8080
    
    spring.jpa.show-sql=true
    spring.jpa.database=mysql
    spring.jpa.database-platform=mysql
    spring.jpa.hibernate.ddl-auto=update
    spring.jpa.properties.hibernate.dialect=org.hibernate.dialect.MySQL57Dialect
   ```

3. **POJO**

   ```java
   import lombok.Data;
   import org.springframework.data.jpa.domain.support.AuditingEntityListener;
   
   import javax.persistence.*;
   
   @Entity
   @Data
   @Table(name = "tb_Customer")
   @EntityListeners(AuditingEntityListener.class)
   public class Customer {
   
       @Id
       @GeneratedValue(strategy = GenerationType.IDENTITY)
       private Long id;
   
       @Column(name = "cust_name")
       private String custName;
   
   }
   ```

4. **Repository/CustomerRepository**

   ```java
    public interface CustomerRepository extends PagingAndSortingRepository<Customer,Long> {
    }
   ```

5. **Service & ServiceImpl**

   ```java
    public interface CustomerService {
        Iterable<Customer> getAll();
    }
   ```

   ```java
   import org.springframework.beans.factory.annotation.Autowired;
   import org.springframework.stereotype.Service;
   
   @Service
   public class CustomerServiceImpl implements CustomerService{
   
       @Autowired
       CustomerRepository repository;
   
       @Override
       public Iterable<Customer> getAll() {
           return repository.findAll();
       }
   }
   ```

6. **Controller**

   ```java
    @RestController
    public class CustomerController {
    
        @Autowired
        CustomerService customerService;
    
        @RequestMapping("/all")
        public Iterable<Customer> getAll() {
            return customerService.getAll();
        }
    }
   ```

   



## For My Reference In Future

[ZhaoStudy Learning Notes on same video](https://www.cnblogs.com/zhaostudy/p/16498332.html)
