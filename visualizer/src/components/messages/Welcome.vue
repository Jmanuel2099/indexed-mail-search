<template>
    <section class="welcome_section">
        <img alt="App_logo" class="logo" src="@/assets/logo.svg" width="125" height="125" />

        <div class="wrapper">
            <div class="greetings">
                <h1 style="color: #0078bd; font-size: x-large; font-weight: bold;">
                    Welcome to the e-mail search engine
                </h1>
                <h3>Your search will be performed with the help of the
                    <a href="https://zincsearch.com/" 
                    target="_blank" rel="noopener"
                    style="color: #0078bd; font-weight: bold;"> Zinc</a>
                    search engine and the 
                    <a href="https://go.dev/" 
                    target="_blank" 
                    rel="noopener"
                    style="color: #0078bd; font-weight: bold;">Go</a>
                    programming language.
                </h3>
            </div>
        </div>
    </section>
</template>

<style scoped>
.welcome_section {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    margin-left: 4rem;
    margin-right: 3rem;
}

.logo {
    display: block;
    margin: 0 auto 2rem;
    background-color: #0078bd;
    border-radius: 1rem;
}

@media (min-width: 1024px) {
    .welcome_section {
        display: flex;
        place-items: center;
        padding-right: calc(var(--section-gap) / 2);
    }

    .logo {
        margin: 0 2rem 0 0;
    }

    .welcome_section .wrapper {
        display: flex;
        place-items: flex-start;
        flex-wrap: wrap;
    }
}

@media (min-width: 1024px) {

    .greetings h1,
    .greetings h3 {
        color: #575656;
        font-family: "Open Sans", sans-serif;
        text-align: left;
    }
}
</style>