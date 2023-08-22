import java.util.Random;

public class HasQuarterState implements State {
  private static final long serialVersionUID = 2L;

  public HasQuarterState(GumballMachine gumballMachine) {
    this.gumballMachine = gumballMachine;
  }

  Random randomWinner = new Random(System.currentTimeMillis());
  transient GumballMachine gumballMachine;

  public void insertQuarter() {
    System.out.println("You can't insert another quarter");
  }

  public void ejectQuarter() {
    System.out.println("Quarter returned");
    gumballMachine.setState(gumballMachine.getNoQuarterState());
  }

  public void turnCrank() {
    System.out.println("You turned...");
    int winner = randomWinner.nextInt(10);
    if ((winner == 0) && (gumballMachine.getCount() > 1)) {
      gumballMachine.setState(gumballMachine.getWinnerQuarterState());
    } else {
      gumballMachine.setState(gumballMachine.getSoldQuarterState());
    }
  }

  public void dispense() {
    System.out.println("No gumball dispensed");
  }

  public void refill() {
  }
}
