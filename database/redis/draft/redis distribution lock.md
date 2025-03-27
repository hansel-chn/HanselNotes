* Some questions and thoughts
>```
>Superficially this works well, but there is a problem: this is a single point of failure in our architecture. What happens if the Redis master goes down? Well, let’s add a replica! And use it if the master is unavailable. This is unfortunately not viable. By doing so we can’t implement our safety property of mutual exclusion, because Redis replication is asynchronous.
>```
>
>https://redis.io/docs/latest/develop/use/patterns/distributed-locks/