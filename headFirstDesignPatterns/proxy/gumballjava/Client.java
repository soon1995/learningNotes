import java.rmi.Naming;

public class Client {

  public static void main(String[] args) {
    String[] locations = {
        "rmi://127.0.0.1:11111/gumballmachine",
        "rmi://127.0.0.1:11112/gumballmachine",
        "rmi://127.0.0.1:11113/gumballmachine",
        // "rmi://aaa.com/gumballmachine",
        // "rmi://bbb.com/gumballmachine",
        // "rmi://ccc.com/gumballmachine"
    };

    GumballMonitor[] monitors = new GumballMonitor[locations.length];

    for (int i = 0; i < locations.length; i++) {
      try {
        System.out.println("1");
        GumballMachineRemote machine = (GumballMachineRemote) Naming.lookup(locations[i]);
        System.out.println("2");
        monitors[i] = new GumballMonitor(machine);
        System.out.println(monitors[i]);
      } catch (Exception e) {
        e.printStackTrace();
      }
    }
  }
}

class GumballMonitor {
  GumballMachineRemote machine;

  public GumballMonitor(GumballMachineRemote machine) {
    this.machine = machine;
  }

  public void report() {
    try {
      System.out.printf("Gumball Machine: %s%n", machine.getLocation());
      System.out.printf("Current inventory: %d gumballs%n", machine.getCount());
      System.out.println("Current State: " + machine.getState());
    } catch (Exception e) {
      e.printStackTrace();
    }
  }
}
