### Configure Auth0

#### Configure Team Access

Using Kore, team IAM (Identity and Access management) [is greatly simplified](security-gke.md#rbac).

Kore uses an external identity provider, like Auth0 or an enterprise's existing SSO system, to directly manage team member access to the team's provisioned environment.

For this guide, we'll be using Auth0 to configure team access.

[Auth0](https://auth0.com/), provides an enterprise SAAS identity provider.

Sign up for an account from the [home page](https://auth0.com).

From the dashboard side menu choose `Applications` and then `Create Application`

Give the application a name and choose `Regular Web Applications`

Once provisioned click on the `Settings` tab and scroll down to `Allowed Callback URLs`.
These are the permitted redirects for the applications. Since we are running the application locally off the laptop set
```
http://localhost:10080/oauth/callback,http://localhost:3000/auth/callback
```

Please make a note of the [__*Domain, Client ID, and Client Secret*__].

Scroll to the bottom of the settings and click the `Show Advanced Settings`

Choose the `OAuth` tab from the advanced settings and ensure that the `JsonWebToken Signature Algorithm` is set to RS256 and `OIDC Conformant` is toggled on.

#### Configuring test users

Return to the Auth0 dashboard. From the side menu select 'Users & Roles' setting.

- Create a user by selecting 'Users'.
- Create a role by selecting 'Roles'.
- Add the role to the user.
