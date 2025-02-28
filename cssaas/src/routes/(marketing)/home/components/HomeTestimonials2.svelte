<script lang="ts">


</script>


<!-- Testimonials -->
<section>
    <div class="max-w-3xl mx-auto px-4 sm:px-6">
        <div class="relative pb-12 md:pb-20">

            <!-- Particles animation -->
            <div class="absolute left-1/2 -translate-x-1/2 top-0 -z-10 w-80 h-80 -mt-6">
                <div class="absolute inset-0 -z-10" aria-hidden="true">
                    <canvas data-particle-animation data-particle-quantity="10" data-particle-staticity="30"></canvas>
                </div>
            </div>
            
            <!-- Carousel -->
             <!-- x-data="carousel" -->
            <div class="text-center" >
                <!-- Testimonial image -->
                <div class="relative h-32 [mask-image:linear-gradient(0deg,transparent,var(--color-white)_40%,var(--color-white))]">
                    <div class="absolute top-0 left-1/2 -translate-x-1/2 w-[480px] h-[480px] -z-10 pointer-events-none before:rounded-full rounded-full before:absolute before:inset-0 before:bg-linear-to-b before:from-slate-400/20 before:to-transparent before:to-20% after:rounded-full after:absolute after:inset-0 after:bg-slate-900 after:m-px before:-z-20 after:-z-20">
                        <!-- Alpine.js template for testimonial images: https://github.com/alpinejs/alpine#x-for -->
                         <!--  x-for="(item, index) in items" :key="index" -->
                        <template>
                            <!-- x-show="active === index" -->
                            <div
                                class="absolute inset-0 -z-10"
                                x-transition:enter="transition ease-[cubic-bezier(0.68,-0.3,0.32,1)] duration-700 order-first"
                                x-transition:enter-start="opacity-0 -rotate-[60deg]"
                                x-transition:enter-end="opacity-100 rotate-0"
                                x-transition:leave="transition ease-[cubic-bezier(0.68,-0.3,0.32,1)] duration-700"
                                x-transition:leave-start="opacity-100 rotate-0"
                                x-transition:leave-end="opacity-0 rotate-[60deg]"
                            >
                                <img class="relative top-11 left-1/2 -translate-x-1/2 rounded-full" :src="item.img" width="56" height="56" :alt="item.name">
                            </div>
                        </template>                                
                    </div>
                </div>
                <!-- Text -->
                <div class="mb-10">
                    <!--  x-ref="testimonials" -->
                    <div class="relative flex flex-col transition-all duration-150 delay-300 ease-in-out">
                        <!-- Alpine.js template for testimonials: https://github.com/alpinejs/alpine#x-for -->
                        <!-- x-for="(item, index) in items" :key="index" -->
                        <template>
                            <!-- x-show="active === index" -->
                            <div
                                x-transition:enter="transition ease-in-out duration-500 delay-200 order-first"
                                x-transition:enter-start="opacity-0 -translate-x-4"
                                x-transition:enter-end="opacity-100 translate-x-0"
                                x-transition:leave="transition ease-out duration-300 delay-300 absolute"
                                x-transition:leave-start="opacity-100 translate-x-0"
                                x-transition:leave-end="opacity-0 translate-x-4"
                            >
                                <!--  x-text="item.quote" -->
                                <div class="text-xl font-bold bg-clip-text text-transparent bg-linear-to-r from-slate-200/60 via-slate-200 to-slate-200/60"></div>
                            </div>
                        </template>
                    </div>
                </div>
                <!-- Buttons -->
                <div class="flex flex-wrap justify-center -m-1.5">
                    <!-- Alpine.js template for buttons: https://github.com/alpinejs/alpine#x-for -->
                    <!-- x-for="(item, index) in items" :key="index" -->
                    <template>
                        <!--
                        :class="active === index ? 'opacity-100' : 'opacity-30 hover:opacity-60'"
                        @click="active = index; stopAutorotate();"
                        -->
                        <button
                            class="btn-sm m-1.5 text-xs py-1.5 text-slate-300 transition duration-150 ease-in-out [background:linear-gradient(var(--color-slate-900),var(--color-slate-900))_padding-box,conic-gradient(var(--color-slate-400),var(--color-slate-700)_25%,var(--color-slate-700)_75%,var(--color-slate-400)_100%)_border-box] relative before:absolute before:inset-0 before:bg-slate-800/30 before:rounded-full before:pointer-events-none"
                            
                        >
                            <span class="relative">
                                <!--  x-text="item.name" -->
                                <span class="text-slate-50"></span> 
                                <span class="text-slate-600">-</span> 
                                <!--  x-text="item.role" -->
                                <span></span>
                            </span>
                        </button>
                    </template>
                </div>
            </div>
            <!-- Carousel data and functionality: https://github.com/alpinejs/alpine -->
            <script>
                document.addEventListener('alpine:init', () => {
                    Alpine.data('carousel', () => ({
                        active: 0,
                        autorotate: true,
                        autorotateTiming: 7000,
                        items: [
                            {
                                img: './images/testimonial-01.jpg',
                                quote: "The ability to capture responses is a game-changer. If a user gets tired of the sign up and leaves, that data is still persisted. Additionally, it's great to be able to select between formats.",
                                name: 'Jessie J',
                                role: 'Ltd Head of Product'
                            },
                            {
                                img: './images/testimonial-02.jpg',
                                quote: "I have been using this product for a few weeks now and I am blown away by the results. My skin looks visibly brighter and smoother, and I have received so many compliments on my complexion.",
                                name: 'Mark Luk',
                                role: 'Spark Founder & CEO'
                            },
                            {
                                img: './images/testimonial-03.jpg',
                                quote: "As a busy professional, I don't have a lot of time to devote to working out. But with this fitness program, I have seen amazing results in just a few short weeks. The workouts are efficient and effective.",
                                name: 'Jeff Kahl',
                                role: 'Appy Product Lead'
                            },
                        ],
                        init() {
                            if (this.autorotate) {
                                this.autorotateInterval = setInterval(() => {
                                    this.active = this.active + 1 === this.items.length ? 0 : this.active + 1
                                }, this.autorotateTiming)
                            }
                            this.$watch('active', callback => this.heightFix())
                        },
                        stopAutorotate() {
                            clearInterval(this.autorotateInterval)
                            this.autorotateInterval = null
                        },
                        heightFix() {
                            this.$nextTick(() => {
                                this.$refs.testimonials.style.height = this.$refs.testimonials.children[this.active + 1].offsetHeight + 'px'
                            })
                        }
                    }))
                })
            </script>
        </div>
    </div>
</section>
