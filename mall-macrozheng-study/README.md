# Macro _ Mall Learning Note

This is a learning notes from [macrozheng/mall](https://github.com/macrozheng/mall)

Start from : 16/7/2022

End: 23/7/2022



## Structure

```
mall
├── mall-common -- 工具类及通用代码
├── mall-mbg -- MyBatisGenerator生成的数据库操作代码
├── mall-security -- SpringSecurity封装公用模块
├── mall-admin -- 后台商城管理系统接口
├── mall-search -- 基于Elasticsearch的商品搜索系统
├── mall-portal -- 前台商城系统接口
└── mall-demo -- 框架搭建时的测试代码
```



## Knowledge Acquired

### `mall-common` Module

Clean Structure

@Builder

Swagger Properties for different Module (used template pattern)

Spring -> AOP to log the controller logs.

> @Aspect
>
> @Component

- Around (Start)
- Before
- Your PointCut
- After
- Around (End)

RedisService to provide a easier code for developer. 

The structure is clean.



### `mall-mbg` Module

handling comment

MyBatis Generator Configuration



### `mall-security` Module

Detail on Jwt Coding

- ```java
  /**
   * JwtToken生成的工具类
   * JWT token的格式：header.payload.signature
   * header的格式（算法、token的类型）：
   * {"alg": "HS512","typ": "JWT"}
   * payload的格式（用户名、创建时间、生成时间）：
   * {"sub":"wang","created":1489079981393,"exp":1489684781}
   * signature的生成算法：
   * HMACSHA512(base64UrlEncode(header) + "." +base64UrlEncode(payload),secret)
   * Created by macro on 2018/4/26.
   */
  ```

- ```java
  //generate token
  Jwts.builder()
      .setClaims(claims)
      .setExpiration(generateExpirationDate())
      .signWith(SignatureAlgorithm.HS512, secret)
      .compact();
  (return token in string)
  ```

- ```java
  //get claim
  private Claims getClaimsFromToken(String token) {
      Claims claims = null;
      try {
          claims = Jwts.parser()
              .setSigningKey(secret)
              .parseClaimsJws(token)
              .getBody();
      } catch (Exception e) {
          LOGGER.info("JWT格式验证失败:{}", token);
      }
      return claims;
  }
  ```

- ```java
  //get username
  //getSubject : claim key = "sub"
  public String getUserNameFromToken(String token) {
      String username;
      try {
          Claims claims = getClaimsFromToken(token);
          username = claims.getSubject();
      } catch (Exception e) {
          username = null;
      }
      return username;
  }
  ```



Able to get bean by using a SpringUtil Class, which save the applicationContext

Register Bean(s) in configuration class

Two ways to get env (the detail in properties file):

- use @Value("${}")
- use @Autowired private Environment environment => environment.get...
- @ConfigurationProperties(prefix = "xx.xx") on class

Save the spring security bean:

- IgnoreUrlsConfig (List`<String>`) 白名单

- restfulAccessDeniedHandler 自定义返回结果：没有权限访问时

- restAuthenticationEntryPoint 自定义返回结果：未登录或登录过期

- jwtAuthenticationTokenFilter JWT登录授权过滤器

- DynamicSecurity's Component : Filter 

  > Dynamic Security 中的有用到的Component

  - Service 在不同module里加载不同的urlPattern
  - MetadataSource 动态权限数据源，用于获取动态权限规则
    - 从数据库拿 urlpattern, id, 和名字，
    - 如果urlPattern和request的path相同，把id和名字存入configAttr list.
    - return Collection<ConfigAttribute> ，代表此path中合格的role
    - 这个供AccessDecisionManager决定 Role(该用户的role)和此 list(该url的role)是否匹配，如果匹配则ok,不匹配抛异常
  - Filter 动态权限过滤器，用于实现基于路径的动态权限过滤 (白名单和 是否Method：OPTION)，过后条用父类的AccessDecisionManager
  - AccessDecisionManager 动态权限决策管理器，用于判断用户是否有访问权限



### `mall-admin` Module

- config
  - BeanPostProcessor => can do something with bean before or after bean initialization
  
    - example here eliminated handlerMappings which do not have patternParser
  
  - SecurityConfig
  
    - @Bean your UserDetailService to get the UserDetails
  
      - > Tips: somewhere must implements UserDetails which override getAuthorities
  
      - can save you DB:ResourceName:UserId - list => into redis as it will be frequently taken out to compare the role.
  
    - @Bean DynamicSecurityService as the resource (permission) might be different in different module
  
  - @Bean CorsFilter 全局跨域配置

- validator

  - ```java
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-validation</artifactId>
    </dependency>
    ```

  - we can set a @interface to set rules on the field.

  - Going to do with 4 classes, which is annotation, validatorClass, entity, and controller

    - Annotation

      ```java
      @Documented
      @Retention(RetentionPolicy.RUNTIME)
      @Target({ElementType.FIELD,ElementType.PARAMETER})
      @Constraint(validatedBy = FlagValidatorClass.class)
      public @interface FlagValidator {
          String[] value() default {};
      
          String message() default "flag is not found";
      
          Class<?>[] groups() default {};
      
          Class<? extends Payload>[] payload() default {};
      }
      
      ```

    - Validator Class

      ```java
      public class FlagValidatorClass implements ConstraintValidator<ValidatorAnnotation,Integer> {
          private String[] values;
          @Override
          public void initialize(ValidatorAnnotation flagValidator) {
              this.values = flagValidator.value();
          }
      
          @Override
          public boolean isValid(Integer value, ConstraintValidatorContext constraintValidatorContext) {
              boolean isValid = false;
              if(value==null){
                  //当状态为空时使用默认值
                  return true;
              }
              for(int i=0;i<values.length;i++){
                  if(values[i].equals(String.valueOf(value))){
                      isValid = true;
                      break;
                  }
              }
              return isValid;
          }
      }
      ```

    - entity

      ```java
      @ValidatorAnnotation(value = {"1","2","3"}, message = "invalid input")
          Integer value;
      ```

    - Controller

      ```java
      @PostMapping("testValidation")
      public String testValidation(@Valid @RequestBody Entity entity, BindingResult result) {
          if (result.hasErrors()) {
              return "false";
          } else {
              return "YESSS";
          }
      }
      ```

      



### `mall-portal` Module

A rabbitMQ configuration for exchange, routing key and queue

A JacksonConfig to prevent sending a field which is null

For RabbitMQ, here was used to cancel order after some duration

- In omsService (Order Management System) generate order, use cancelOrderSender (a class to send message to TTL Queue)
- In CancelOrderReceiver, call the function delete
- The flow : generating => send message in TTL queue => once time due, cancelOrder

A spring Scheduling to cancelTimeOutOrder every 10 minutes

Transaction Recap

- Read-uncommited ==> Dirty read, Non repeatable read, Phantom read

- Read-commited ==> Non repeatable read, Phantom read

- Repeatable-Read ==> Phantom read

- Serializable ==> n/a

- To use spring transaction

  - Dependency (normally inclusive in jdbc)
  - @EnableTransactionManagement <= in config
  - @Transaction for methods

- @**Transactional**注解失效的情况 (refer [光速膨胀的鹏鹏](https://www.zhihu.com/people/pisecespeng))

  > 1. 引擎是否支持事务, 例如:**mysql数据库表引擎InnoDB支持事务, 但MyISAM不支持**;
  >
  > 2. @Transactional注解的方法需要public修饰;
  >
  > 3. @Transactional注解所在的类不由**Spring容器**管理;
  >
  > 4. 注意是否check异常(**@Transactional注解**的'rollbackFor'和'noRollbackFor');
  >
  > 5. **不带事务的方法调用同一个类中带事务的方法, 事务不会生效;**
  >
  >    eg case when A() {.... B()};
  >
  >    |                      | A()                 | @Transaction A() |
  >    | -------------------- | ------------------- | ---------------- |
  >    | **B()**              | no transaction work | transaction work |
  >    | **@Transaction B()** | no transaction work | transaction work |

- The used of mongodb to record the, majorly used for member_bookmarkBrand / productHistory/ productType ids operation

  
