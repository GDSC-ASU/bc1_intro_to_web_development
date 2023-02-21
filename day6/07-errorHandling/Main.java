class Main {
  public static double sqrt(int n) throws Exception {
    if (n < 0) {
      throw new NegativeNumberException;
    }
    return Math.sqrt(n);
  }

  public static void main(String[] args) {
    try {
      double value = sqrt(-12);
    } catch (Exception err) {
      System.out.println(err.getMessage());
    }
  }
}
