<script lang="ts">
  let name: string | undefined;
  let age: number | undefined;
  let email: string | undefined;
  let description: string | undefined;
  const createPost = async (e: SubmitEvent) => {
    const url = `http://localhost:${import.meta.env.VITE_POL_BACK_PORT}`;
    console.log(url);

    try {
        const res = await fetch(url + "/post", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                "name": name,
                "age": age,
                "email": email,
                "description": description
            })
        });

        if (!res.ok) {
            throw new Error(`HTTP error! status: ${res.status}`);
        }
        const data = await res.json();
        console.log(data);
    } catch (error) {
        console.error('Error:', error);
    }
  }
</script>

<div class="form">
    <form on:submit|preventDefault={createPost}>
        <div class="form-inputs">
            <label for="name">
                <input type="text" id="name" placeholder="Mohammed" bind:value={name}/>
            </label>
            <label for="age">
                <input type="number" id="age" placeholder="19"  bind:value={age}/>
            </label>
            <label for="email">
                <input type="text" id="email" placeholder="mohammed@test.com" bind:value={email}/>
            </label>
            <label for="email">
                <input type="text" id="description" placeholder="Hello World!"  bind:value={description}/>
            </label>
        </div>
        <button id="post-button">
          Submit
        </button>
    </form>
</div>

<style>
    .form {
        margin-top: 15rem;
        margin-left: 45rem;
        width: 400px;
        height: 400px;
        background-color: rgb(196, 193, 188);
        border-radius: 15px;
        box-shadow: 5px 0 5px 0;
    }

    .form-inputs {
        padding-top: 5rem;
        margin-top: 5rem;
        margin-left: 6rem;
    }

    .form input {
        justify-content: center;
        margin-top: 5px;
        border-radius: 5px;
	    color: inherit;
	    border: none;
	    font: inherit;
	    outline: inherit;
        width: 200px;
        height: 35px;
    }
    .form button {
        margin-top: 5rem;
        margin-left: 9rem;
        color: white;
        border: none;
        background-color: #151313;
        border-radius: 5px;
        width: 100px;
        height: 30px;
        cursor: pointer;
    }
</style>
