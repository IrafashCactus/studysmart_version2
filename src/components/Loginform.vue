<template>
  <div class="flex items-center justify-center p-6">
    <Form @submit="submitForm" class="w-full max-w-md rounded-3xl bg-white p-8 shadow-2xl ring-1 ring-slate-900/10">
      <div class="mb-6 text-center">
        <h2 class="text-3xl font-semibold text-slate-900">Welcome back</h2>
        <p class="mt-2 text-sm text-slate-500">Login to access your personalized study dashboard.</p>
      </div>

      <FormField name="username" label="Username" class="flex flex-col gap-2">
        <InputText id="home-login-username" name="username" v-model="username" type="text" placeholder="Username" autocomplete="username" class="w-full" />
        <Message
          v-if="submitted && !username"
          severity="error"
          size="small"
          variant="simple"
          text="Username is required."
        />
      </FormField>

      <FormField name="password" label="Password" class="flex flex-col gap-2">
        <Password
          id="home-login-password"
          name="password"
          v-model="password"
          placeholder="Password"
          autocomplete="current-password"
          :feedback="false"
          toggleMask
          class="w-full"
        />
        <Message
          v-if="submitted && !password"
          severity="error"
          size="small"
          variant="simple"
          text="Password is required."
        />
      </FormField>

      <Button type="submit" severity="secondary" label="Login" class="w-full rounded-xl py-3 text-base font-semibold" />

      <p v-if="loginSuccess" class="mt-4 rounded-2xl bg-emerald-50 px-4 py-3 text-sm text-emerald-700">{{ loginSuccess }}</p>

      <Message
        v-if="loginError"
        severity="error"
        size="small"
        variant="simple"
        :text="loginError"
      />
    </Form>
  </div>
</template>


<script lang="js">
import { Form } from '@primevue/forms'
import FormField from '@primevue/forms/formfield'
import InputText from 'primevue/inputtext'
import Password from 'primevue/password'
import Button from 'primevue/button'
import Message from 'primevue/message'
import auth from '@/services/auth'

export default {
  name: 'LoginForm',
  components: {
    Form,
    FormField,
    InputText,
    Password,
    Button,
    Message,
  },
  data() {
    return {
      username: '',
      password: '',
      submitted: false,
      loginError: '',
      loginSuccess: '',
    }
  },
  methods: {
    async submitForm(e) {
      if (e && typeof e.preventDefault === 'function') e.preventDefault()
      this.submitted = true
      this.loginError = ''
      this.loginSuccess = ''

      if (!this.username || !this.password) {
        this.loginError = 'Please enter both username and password.'
        return
      }

      try {
        await auth.login({
          username: this.username,
          password: this.password,
        })

        this.loginSuccess = 'Login successful! Welcome back.'
        this.$router.push('/')
      } catch (error) {
        console.error('Login error:', error)
        this.loginError =
          error?.response?.data?.message ||
          error?.message ||
          'Login failed. Please try again.'
      }
    },
  },
}
</script>