# 💡Overview

Imagine you’re building a rewards system for a loyalty program. To prevent abuse, the system should enforce limits on how frequently a user can redeem rewards. For instance, a user might only be allowed to redeem up to **N** rewards within **T** seconds.

Your task is to implement a component that enforces this redemption limit **per user**.

## **✳️ Requirements**

- The limiter should **track reward redemptions per user**.
- Each user should be allowed to redeem up to **limit** rewards within a sliding or fixed **time window** of **T** seconds.
- Any redemption requests that exceed this limit within the current window should be **rejected**.
- Rejected requests should **not affect** the count or timing of future accepted requests.
- The solution should work **in-memory** (no external storage or services).

### **✅ Example Behavior**

If the limit is 3 rewards per 10 seconds:

```shell
User A redeems at t=0s   ✅ allowed  
User A redeems at t=2s   ✅ allowed  
User A redeems at t=5s   ✅ allowed  
User A redeems at t=7s   ❌ rejected  
User A redeems at t=11s  ✅ allowed (t=0s has expired)
```

## **📝 Instructions**

Please follow these steps to complete the challenge:

1.	Clone the Repository
```shell
git clone https://github.com/your-org/repo-name.git
cd repo-name
```

2.	Implement the Solution
	
3.	Push Your Code

When you’re finished, push your branch to GitHub:

```shell
git add .
git commit -m 'feat: Final implementation'
git push
```

4. Submit a PR against the source repo

There should be a branch on the source repo with your first name + last initial. Please submit a PR against that branch.
