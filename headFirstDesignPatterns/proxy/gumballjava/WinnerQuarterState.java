public class WinnerQuarterState implements State {
  private static final long serialVersionUID = 2L;

  transient GumballMachine gumballMachine;

  public WinnerQuarterState(GumballMachine gumballMachine) {
    this.gumballMachine = gumballMachine;
  }

  public void insertQuarter() {
    System.out.println("Please wait, we're already giving you a gumball");
  }

  public void ejectQuarter() {
    System.out.println("Sorry, you already turned the crank");
  }

  public void turnCrank() {
    System.out.println("Turning twice doesn't get you another gumball!");
  }

  public void dispense() {
    gumballMachine.ReleaseBall();
    if (gumballMachine.getCount() == 0) {
      gumballMachine.setState(gumballMachine.getSoldOutQuarterState());
    } else {
      gumballMachine.ReleaseBall();
      System.out.println("YOU'RE A WINNER! You got two gumballs for your quarter");
      if (gumballMachine.getCount() > 0) {
        gumballMachine.setState(gumballMachine.getNoQuarterState());
      } else {
        System.out.println("Oops, out of gumballs!");
        gumballMachine.setState(gumballMachine.getSoldOutQuarterState());
      }
    }
  }

  public void refill() {
  }
}
