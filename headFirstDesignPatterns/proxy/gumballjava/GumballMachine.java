import java.rmi.RemoteException;
import java.rmi.server.UnicastRemoteObject;

public class GumballMachine extends UnicastRemoteObject implements GumballMachineRemote {

  private static final long serialVersionUID = 2L;

  private int count;
  private String location;

  private State noQuarterState;
  private State hasQuarterState;
  private State soldQuarterState;
  private State soldOutQuarterState;
  private State winnerQuarterState;

  private State state;

  public GumballMachine(String location, int count) throws RemoteException {
    noQuarterState = new NoQuarterState(this);
    hasQuarterState = new HasQuarterState(this);
    soldQuarterState = new SoldQuarterState(this);
    soldOutQuarterState = new SoldOutQuarterState(this);
    winnerQuarterState = new WinnerQuarterState(this);
    this.location = location;
    this.count = count;
    if (count > 0) {
      state = noQuarterState;
    } else {
      state = soldOutQuarterState;
    }
  }

  public void setState(State state) {
    this.state = state;
  }

  public int getCount() {
    return count;
  }

  public void setCount(int count) {
    this.count = count;
  }

  public String getLocation() {
    return location;
  }

  public State getNoQuarterState() {
    return noQuarterState;
  }

  public State getHasQuarterState() {
    return hasQuarterState;
  }

  public State getSoldQuarterState() {
    return soldQuarterState;
  }

  public State getSoldOutQuarterState() {
    return soldOutQuarterState;
  }

  public State getWinnerQuarterState() {
    return winnerQuarterState;
  }

  public State getState() {
    return state;
  }

  public void ReleaseBall() {
    System.out.println("A gumball comes rolling out the slot...");
    if (count > 0) {
      count--;
    }
  }

  public void Refill(int count) {
    count += count;
    System.out.printf("The gumball machine was just refilled; its new count is : %d%n", count);
    state.refill();
  }
}
