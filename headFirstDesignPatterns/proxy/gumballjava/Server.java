import java.rmi.Naming;
import java.rmi.registry.LocateRegistry;
import java.rmi.registry.Registry;

public class Server {
  public static void main(String[] args) {
    GumballMachineRemote gumballMachine = null;
    int count;

    if (args.length < 2) {
      // System.out.println("GumballMachine <name> <inventory>");
      System.out.println("GumballMachine <port> <inventory>");
      System.exit(1);
    }

    try {
      count = Integer.parseInt(args[1]);
      gumballMachine = new GumballMachine(args[0], count);

      int port = Integer.parseInt(args[0]);
      Registry registry = LocateRegistry.createRegistry(port);
      registry.rebind("gumballmachine", gumballMachine);
      // Naming.rebind("//" + args[0] + "/gumballmachine", gumballMachine);
    } catch (Exception e) {
      e.printStackTrace();
    }
  }
}
