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
        <input type="text" id="name" placeholder="Mohammed" bind:value={name}/>
        <label for="name"/>
        <input type="number" id="age" placeholder="19"  bind:value={age}/>
        <label for="age"/>
        <input type="text" id="email" placeholder="mohammed@test.com" bind:value={email}/>
        <label for="email"/>
        <input type="text" placeholder="Hello World!"  bind:value={description}/>
        <label for="email"/>
        <button>
          Submit
        </button>
    </form>
</div>

<style>
    .form {
        margin-left: 50%;
        width: 400px;
        height: 400px;
        background-color: rgb(119, 209, 128);
        border-radius: 15px;
        display: flex;
    }
    .form input {
      border-radius: 5px;
	    color: inherit;
	    border: none;
	    font: inherit;
	    outline: inherit;
      width: 200px;
      height: 35px;
    }
    .form button {
	    border: none;
      background-color: red;
      border-radius: 5px;
      width: 80px;
      height: 25px;
    }
</style>
