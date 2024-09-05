<script lang="ts">
    import { parseErrors, mergeErrors } from '$lib/forms/helpers'
    import { loginForm, loginSchema } from '$lib/forms/login.validation'
    import { login } from '$lib/services/auth'
    import { goto } from '$app/navigation'
    import { EmailInput, PasswordInput } from "$lib/components/forms"
    import { AlertError } from "$lib/components/messages"

    export let errors: any = {}
    export let lf = loginForm

    let loginFailed:Boolean = false

    const submitForm = async () => {
		errors = {}

		const bs = loginSchema.validate(lf, { abortEarly: false })
			.then(() => {
				console.log('validation passed')

                login({ email: lf.email, password: lf.password })
                    .then((r) => {
                        console.log(r.data.login.status?.success.valueOf())
                        if(r.data.login.status?.success.valueOf()) {
                            loginFailed = false
                            goto('/home')
                        } else {
                            loginFailed = true
                        }
                    })
			})
			.catch((err) => {
				errors = mergeErrors(errors, parseErrors(err))

				console.log(errors);
			});
    }

    $: lf
	$: errors
    $: loginFailed
</script>
  

<div class="mt-8">
<div class="w-full p-4 bg-white border border-gray-200 rounded-lg shadow sm:p-6 md:p-8 dark:bg-gray-800 dark:border-gray-700">
<form class="space-y-6" action="#">
<h5 class="text-xl font-medium text-gray-900 dark:text-white">Sign in to Analyzer</h5>
    <AlertError showAlert={loginFailed} title="Login Attempt Failed" message="We were unable to validate your login credentials"/>
    
    <EmailInput bind:value={lf.email} 
        fieldName="Email" 
        error={errors.email} />
    
    <PasswordInput bind:value={lf.password} 
        fieldName="Password" 
        error={errors.password} />

    <button type="button" on:click={submitForm} class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Login</button>
    <div class="text-sm font-medium text-gray-500 dark:text-gray-300">
        Not registered? <a href="/register" class="text-blue-700 hover:underline dark:text-blue-500">Create account</a> <br />
        Forgot password? <a href="/forgot" class="text-blue-700 hover:underline dark:text-blue-500">Reset</a>
    </div>
</form>
</div>
</div>