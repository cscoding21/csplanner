<script lang="ts">
    import ai from '$lib/assets/images/auth-illustration.svg'
    import logoImage from '$lib/assets/images/csplanner-logo-white.png?enhanced'

</script>


<section class="relative">
    
    <!-- Illustration -->
    <div class="md:block absolute left-1/2 -translate-x-1/2 -mt-36 blur-2xl opacity-70 pointer-events-none -z-10" aria-hidden="true">
        <img src={ai} class="max-w-none" width="1440" height="450" alt="Page Illustration">
    </div>

    <div class="relative max-w-6xl mx-auto px-4 sm:px-6">
        <div class="pt-32 pb-12 md:pt-40 md:pb-20">

            <!-- Page header -->
            <div class="max-w-3xl mx-auto text-center pb-12">
                <!-- Logo -->
                <div class="mb-5">
                    <a class="inline-flex" href="/home" aria-label="csPlanner Logo">
                        <div class="relative flex items-center justify-center">
                            <enhanced:img class="relative" src={logoImage} width="121" alt="Stellar">
                        </div>
                    </a>
                </div>
                <!-- Page title -->
                <h1 class="h2 bg-clip-text text-transparent bg-linear-to-r from-slate-200/60 via-slate-200 to-slate-200/60">Reset your password</h1>
            </div>

            <!-- Form -->
            <div class="max-w-sm mx-auto">

                <form>
                    <div class="space-y-4">
                        <div>
                            <label class="block text-sm text-slate-300 font-medium mb-1" for="email">Email</label>
                            <input id="email" class="form-input w-full" type="email" required />
                        </div>
                    </div>
                    <div class="mt-6">
                        <button class="btn text-sm text-white bg-purple-500 hover:bg-purple-600 w-full shadow-xs group">
                            Reset Password <span class="tracking-normal text-purple-300 group-hover:translate-x-0.5 transition-transform duration-150 ease-in-out ml-1">-&gt;</span>
                        </button>
                    </div>
                </form>

            </div>

        </div>
    </div>

</section>