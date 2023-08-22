import java.rmi.RemoteException;
import java.rmi.server.UnicastRemoteObject;

public class MyRemoteImpl extends UnicastRemoteObject implements MyRemote {
  // UnicastRemoteObject implements Serializable, so we need the serialVersionUID
  // field
  private static final long serialVersionUID = 1L;

  // we do this because UnicastRemoteObject constructor throws RemoteException
  public MyRemoteImpl() throws RemoteException {
  }

  public String sayHello() {
    return "Server says, 'Hey'";
  }
  // more code in class
}
