const CraftLuaBridgeClient = {
    endpoint: "/bridge",
    send(payload) {
        return fetch(this.endpoint, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(payload)
        });
    }
};