
public class InputTest {
  public static void main(String[] args) throws IOException {
    int c;

    try {
      InputStream in =
        new LowerCaseInputStream(
          new BufferedInputTream
            newFileInputStream("test.txt")));

      while ((c = in.read) >= 0 ) {
        System.out.print((char)c);
      }

      in.close();
    } catch (IOException e) {
      e.printStackTrace();
    }
  }
}
