# Blind SQL injection with time delays and information retrieval

In this lab, we observe that the application does not return any visible response to traditional SQL injection techniques, but we discover that it is possible to trigger time delays.

![Screenshot1](../../04-Screenshots/blind1.png)

To start a new statement separated from the previous one, we use the delimiter `;`, URL-encoded as `%3b`.

![Screenshot2](../../04-Screenshots/blind2.png)

We then test conditional time delays.

![Screenshot3](../../04-Screenshots/blind3.png)

Next, we perform an injection to verify the user and determine the length of the password.

![Screenshot4](../../04-Screenshots/blind4.png)

Finally, we construct the payload that will allow us to retrieve the administrator’s credentials.

We send the payload to Intruder and iterate, brute-forcing each index of the password.

![Screenshot5](../../04-Screenshots/blind5.png)
