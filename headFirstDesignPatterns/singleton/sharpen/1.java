public class Chocolate Boiler {
  private boolean empty;
  private boolean boiled;
  private static Chocolate chocolate
    private ChocolateBoiler() {
      empty = true;
      boiled = false;
    }

  public static Chocolate getInstance() {
    if (chocolate == null) {
      chocolate = new Chocolate()
    }
    return chocolate
  }

  ...
}
