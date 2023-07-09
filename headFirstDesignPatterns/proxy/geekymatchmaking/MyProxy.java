import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Proxy;

public class MyProxy {
  // Person getOwnerProxy(Person person) {
  // return (Person) Proxy.newProxyInstance(
  // person.getClass().getClassLoader(),
  // person.getClass().getInterfaces(),
  // new OwnerInvocationHandler(person));
  // }

  // Person getNonOwnerProxy(Person person) {
  // return (Person) Proxy.newProxyInstance(
  // person.getClass().getClassLoader(),
  // person.getClass().getInterfaces(),
  // new NonOwnerInvocationHandler(person));
  // }

  Person getOwnerProxy(Person person) {
    return getProxy(person, new OwnerInvocationHandler(person));
  }

  Person getNonOwnerProxy(Person person) {
    return getProxy(person, new NonOwnerInvocationHandler(person));
  }

  Person getProxy(Person person, InvocationHandler handler) {
    return (Person) Proxy.newProxyInstance(
        person.getClass().getClassLoader(),
        person.getClass().getInterfaces(),
        handler);
  }
}
