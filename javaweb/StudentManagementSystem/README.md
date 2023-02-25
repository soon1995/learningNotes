# JavaWeb - Student Management System

> Installation of tomcat

1. I referred to [here](https://www.hostinger.my/tutorials/how-to-install-tomcat-on-ubuntu/)

2. in YOUR_INSTALLATION_PATH/conf/Catalina/localhost/, create one xml file, the name does not matter

   ```xml
   <Context path="/myapp" docBase="/yourwebapppath/webapp"/>
   ```

3. To start tomcat: YOUR_INSTALLATION_PATH/bin/startup.sh

4. To stop tomcat: YOUR_INSTALLATION_PATH/bin/shutdown.sh



project structure

```
.           
├── main    
│   ├── java
│   │   └── com           
│   │       └── alsoon    
│   │           └── studentms  
│   │               ├── bean   
│   │               │   ├── User.java            
│   │               │   ├── UserCourse.java      
│   │               │   └── UserEditVo.java      
│   │               ├── dao    
│   │               │   ├── UserCourseDao.java   
│   │               │   ├── UserDao.java         
│   │               │   └── impl                 
│   │               │       ├── UserCourseDaoImpl.java         
│   │               │       └── UserDaoImpl.java 
│   │               ├── service
│   │               │   ├── UserCourseService.java             
│   │               │   ├── UserService.java     
│   │               │   └── impl                 
│   │               │       ├── UserCourseServiceImpl.java     
│   │               │       └── UserServiceImpl.java           
│   │               ├── servlet
│   │               │   ├── CheckScoreServlet.java             
│   │               │   ├── DashboardServlet.java
│   │               │   ├── DeleteUser.java      
│   │               │   ├── EditServlet.java     
│   │               │   ├── LoginServlet.java    
│   │               │   ├── SearchUserServlet.java             
│   │               │   └── ShowEditServlet.java 
│   │               └── util   
│   │ 					├── ConnectionFactory.java             
│   │ 					├── UserIsDeleted.java   
│   │ 					└── UserRole.java        
│   ├── resources         
│   │   └── c3p0-config.xml    
│   └── webapp            
│       ├── WEB-INF       
│       │   └── web.xml   
│       ├── dashboard.jsp 
│       ├── edit.jsp      
│       ├── index.jsp     
│       ├── login.html    
│       ├── login.jsp     
│       ├── logout.jsp    
│       └── score.jsp     
└── test        
	└── java        
		└── com         
			└── soon        
				└── test        
					└── Test.java         

```



pom.xml

```xml
<dependencies>
    <dependency>
      <groupId>junit</groupId>
      <artifactId>junit</artifactId>
      <version>4.11</version>
      <scope>test</scope>
    </dependency>
    <dependency>
      <groupId>javax.servlet</groupId>
      <artifactId>javax.servlet-api</artifactId>
      <version>4.0.1</version>
      <scope>provided</scope>
    </dependency>
    <dependency>
      <groupId>org.projectlombok</groupId>
      <artifactId>lombok</artifactId>
      <version>1.18.20</version>
      <scope>provided</scope>
    </dependency>
    <dependency>
      <groupId>com.mchange</groupId>
      <artifactId>c3p0</artifactId>
      <version>0.9.5.5</version>
    </dependency>
    <dependency>
      <groupId>jstl</groupId>
      <artifactId>jstl</artifactId>
      <version>1.2</version>
    </dependency>
    <dependency>
      <groupId>mysql</groupId>
      <artifactId>mysql-connector-java</artifactId>
      <version>8.0.31</version>
    </dependency>
    <dependency>
      <groupId>commons-beanutils</groupId>
      <artifactId>commons-beanutils</artifactId>
      <version>1.9.4</version>
    </dependency>
  </dependencies>
```

```xml
<plugin>
    <groupId>org.apache.tomcat.maven</groupId>
    <artifactId>tomcat7-maven-plugin</artifactId>
    <version>2.2</version>
    <configuration>
        <port>9090</port>
	</configuration> 
</plugin>	
```

**Command:**

`mvn tomcat7:run`



web.xml

```xml
use this to enable jsp ${} EL 
-----------------------------------
<web-app xmlns="http://java.sun.com/xml/ns/j2ee"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xsi:schemaLocation="http://java.sun.com/xml/ns/j2ee http://java.sun.com/xml/ns/j2ee/web-app_2_4.xsd"
         version="2.4"> </web-app>
```



resources/c3p0-config.xml

```xml
<?xml version="1.0" encoding="UTF-8"?>
<c3p0-config>
	<default-config>
      <property name="driverClass">com.mysql.cj.jdbc.Driver</property>
      <property name="jdbcUrl">jdbc:mysql://localhost:3306/studentms?serverTimezone=UTC</property>
      <property name="user">wsl_root</property>
      <property name="password">password</property>
      <property name="acquireIncrement">10</property>
      <property name="initialPoolSize">10</property>
      <property name="maxPoolSize">100</property>
      <property name="maxIdleTime">60</property>
      <property name="minPoolSize">5</property>
	</default-config>
</c3p0-config>
```



servlet.java

Example 1: 

> - @WebServlet
> - req.getParameter(String)
> - req.getSession() -> session.setAttribute(name, value);
> - req.setAttribute(name, value);
> - req.getRequestDispatcher("/dashboard").forward(req, resp)

```java
package com.alsoon.studentms.servlet;

import java.io.IOException;
import java.util.ArrayList;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;

import com.alsoon.studentms.bean.User;
import com.alsoon.studentms.service.UserService;
import com.alsoon.studentms.service.impl.UserServiceImpl;

@WebServlet("/login")
public class LoginServlet extends HttpServlet {
  UserService userService = new UserServiceImpl();

  @Override
  protected void doPost(HttpServletRequest req, HttpServletResponse resp) throws ServletException, IOException {
    // not doing error handling
    Integer uid = Integer.parseInt(req.getParameter("u_id"));
    String upwd = req.getParameter("u_pwd");
    User user = new User();
    user.setU_id(uid);
    user.setU_pwd(upwd);
    User dbUser = userService.login(user);
    if (dbUser != null && dbUser.getU_id().equals(uid) && dbUser.getU_pwd().equals(upwd)) {
      // set Session Attribute
      HttpSession session = req.getSession();
      session.setAttribute("user", dbUser);
        
      ArrayList<User> students = userService.getAllStudents();
      dbUser.setU_pwd(null);
      req.setAttribute("students", students);

      req.getRequestDispatcher("dashboard.jsp").forward(req, resp);
    } else {
      // relogin
        ...
    }
  }
}

```

Example 2

> - BeanUtils.populate(object, req.getParameterMap()); (UserEditVo has String[] field to receive array)
> - resp.setStatus(val)

```java
@WebServlet("/edit")
public class EditServlet extends HttpServlet{
  UserCourseService userCourseService = new UserCourseServiceImpl();
  @Override
  protected void doPost(HttpServletRequest req, HttpServletResponse resp) throws ServletException, IOException {
    UserEditVo userEditVo = new UserEditVo();
    try {
      BeanUtils.populate(userEditVo, req.getParameterMap());
      userCourseService.updateUserCourseById(userEditVo);
    } catch (IllegalAccessException | InvocationTargetException e) {
      e.printStackTrace();
    }
    req.setAttribute("message", "Successfully updated");
    resp.setStatus(400);
    req.getRequestDispatcher("/dashboard").forward(req, resp);
  }
  
}
```



util/ConnectionFactory

```java
package com.alsoon.studentms.util;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;

import javax.sql.DataSource;

import com.mchange.v2.c3p0.ComboPooledDataSource;

public class ConnectionFactory {
  
  public static DataSource dataSource = new ComboPooledDataSource();

  public static Connection getConnection() {
    try {
      return dataSource.getConnection();
    } catch (SQLException e) {
      e.printStackTrace();
    }
    return null;
  }

  public static void close(Connection connection, PreparedStatement preparedStatement, ResultSet resultSet) {
    try {
      if(resultSet != null) resultSet.close();
      if(preparedStatement != null) preparedStatement.close();
      if(connection != null) connection.close();
    } catch (SQLException e) {
      e.printStackTrace();
    }
  }
}

```

