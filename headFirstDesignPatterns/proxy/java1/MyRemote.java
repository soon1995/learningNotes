import java.rmi.*;

public interface MyRemote extends Remote {
  // Besure arguements and return values are primitives or Serializable
  public String sayHello() throws RemoteException;
}
