# Test results

```shell
=== RUN   Test_htmlTable
=== RUN   Test_htmlTable/maxcian
[Test 1] -->
Given: <table><tr><td>1</td><td>TWO</td></tr><tr><td>three</td><td>FoUr4</td></tr></table>
===== start =====
Target row: 0 - split by <tr> we got 2 row(s):
         0 <td>1</td><td>TWO</td></tr>
         1 <td>three</td><td>FoUr4</td></tr></table>
Target column: 1 - split by <td> we got 2 column(s):
         0 1</td>
         1 TWO</td></tr>
=====  end  =====

[Test 2] -->
Given: <table><tr><td>1</td><td>TWO</td></tr></table>
===== start =====
Target row: 1 - split by <tr> we got 1 row(s):
         0 <td>1</td><td>TWO</td></tr></table>
=====  end  =====

[Test 3] -->
Given: <table><tr><td>7BusWMJ</td><td>D</td><td>5QPh9o</td></tr><tr><td>4Z</td><td>9Z</td><td>okc3</td></tr><tr><td>7mV88j</td><td>K358zV8</td><td>Y2vE</td></tr></table>
===== start =====
Target row: 1 - split by <tr> we got 3 row(s):
         0 <td>7BusWMJ</td><td>D</td><td>5QPh9o</td></tr>
         1 <td>4Z</td><td>9Z</td><td>okc3</td></tr>
         2 <td>7mV88j</td><td>K358zV8</td><td>Y2vE</td></tr></table>
Target column: 1 - split by <td> we got 3 column(s):
         0 4Z</td>
         1 9Z</td>
         2 okc3</td></tr>
=====  end  =====

[Test 7] -->
Given: <table><tr><th>CIRCUMFERENCE</th><th>1</th><th>2</th><th>3</th><th>4</th><th>5</th><th>6</th></tr><tr><td>BITS</td><td>3</td><td>4</td><td>8</td><td>10</td><td>12</td><td>15</td></tr></table>
===== start =====
Target row: 0 - split by <tr> we got 2 row(s):
         0 <th>CIRCUMFERENCE</th><th>1</th><th>2</th><th>3</th><th>4</th><th>5</th><th>6</th></tr>
         1 <td>BITS</td><td>3</td><td>4</td><td>8</td><td>10</td><td>12</td><td>15</td></tr></table>
Target column: 6 - split by <td> we got 0 column(s):
=====  end  =====
```