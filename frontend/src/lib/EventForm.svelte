<script lang="ts">
    import { models } from "$lib/wailsjs/go/models";
    import { getModalStore } from "@skeletonlabs/skeleton";
    const modalStore = getModalStore();
    export let parent: any;
    export let event: models.Event;

    var formEventName: string = event.Name;
    var formStartTime: Date = event.StartTime;
    var formIsRecurring: boolean = event.IsRecurring;

    const TIMEOUT_DURATION_SECOND_FACTOR = 1000000000;
    const TIMEOUT_DURATION_MINUTE_FACTOR = 60 * TIMEOUT_DURATION_SECOND_FACTOR;
    const TIMEOUT_DURATION_HOUR_FACTOR = 60 * TIMEOUT_DURATION_MINUTE_FACTOR;
    const TIMEOUT_DURATION_DAY_FACTOR = 24 * TIMEOUT_DURATION_HOUR_FACTOR;
    var timeoutRemaining: number = event.TimeoutDuration;
    var formTimeoutDurationDays: number = Math.floor(timeoutRemaining / TIMEOUT_DURATION_DAY_FACTOR);
    timeoutRemaining = timeoutRemaining % TIMEOUT_DURATION_DAY_FACTOR;
    var formTimeoutDurationHours: number = Math.floor(timeoutRemaining / TIMEOUT_DURATION_HOUR_FACTOR);
    timeoutRemaining = timeoutRemaining % TIMEOUT_DURATION_HOUR_FACTOR;
    var formTimeoutDurationMinutes: number = Math.floor(timeoutRemaining / TIMEOUT_DURATION_MINUTE_FACTOR);
    timeoutRemaining = timeoutRemaining % TIMEOUT_DURATION_MINUTE_FACTOR;
    var formTimeoutDurationSeconds: number = Math.floor(timeoutRemaining / TIMEOUT_DURATION_SECOND_FACTOR);

    function editEvent() {
        var editedEvent = new models.Event({
            ID: event.ID,
            Name: formEventName,
            StartTime: new Date(formStartTime),
            TimeoutDuration:
                formTimeoutDurationDays * TIMEOUT_DURATION_DAY_FACTOR +
                formTimeoutDurationHours * TIMEOUT_DURATION_HOUR_FACTOR +
                formTimeoutDurationMinutes * TIMEOUT_DURATION_MINUTE_FACTOR +
                formTimeoutDurationSeconds * TIMEOUT_DURATION_SECOND_FACTOR,
            IsRecurring: formIsRecurring,
        });
        if ($modalStore[0] && $modalStore[0].response) $modalStore[0].response(editedEvent);
        modalStore.close();
    }

    function cancelEditing() {
        if ($modalStore[0] && $modalStore[0].response) $modalStore[0].response(undefined);
        modalStore.close();
    }
    parent.onClose = cancelEditing;

    // Base Classes
    const cBase = "card p-4 w-modal shadow-xl space-y-4";
    const cHeader = "text-2xl font-bold";
    const cForm = "border border-surface-500 p-4 space-y-4 rounded-container-token";
</script>

{#if $modalStore[0]}
    <div class={cBase}>
        <header class={cHeader}>{$modalStore[0].title ?? ""}</header>
        <article>{$modalStore[0].body ?? ""}</article>
        <form id="eventForm" class="modal-form {cForm}" on:submit|preventDefault={editEvent}>
            <label class="label">
                <span>Event Name</span>
                <input type="text" class="input" bind:value={formEventName} />
            </label>

            <label>
                <span>Start Time</span>
                <input type="datetime-local" class="input" bind:value={formStartTime} />
            </label>

            <div>
                <span>Timeout Duration</span>
                <div id="TimeoutDurationGrid">
                    <div class="input-group input-group-divider grid-cols-[auto_1fr]">
                        <div class="input-group-shim">Days</div>
                        <input type="number" min=0 bind:value={formTimeoutDurationDays} />
                    </div>
                    <div class="input-group input-group-divider grid-cols-[auto_1fr]">
                        <div class="input-group-shim">Hours</div>
                        <input type="number" min=0 bind:value={formTimeoutDurationHours} />
                    </div>
                    <div class="input-group input-group-divider grid-cols-[auto_1fr]">
                        <div class="input-group-shim">Minutes</div>
                        <input type="number" min=0 bind:value={formTimeoutDurationMinutes} />
                    </div>
                    <div class="input-group input-group-divider grid-cols-[auto_1fr]">
                        <div class="input-group-shim">Seconds</div>
                        <input type="number" min=0 bind:value={formTimeoutDurationSeconds} />
                    </div>
                </div>
            </div>

            <label>
                <span>Is Recurring</span>
                <input type="checkbox" class="input" bind:checked={formIsRecurring} />
            </label>
        </form>
        <footer class="modal-footer {parent.regionFooter}">
            <button class="btn {parent.buttonNeutral}" on:click={cancelEditing}>Cancel</button>
            <button class="btn {parent.buttonPositive}" form="eventForm">Submit Data</button>
        </footer>
    </div>
{/if}

<style>
    #TimeoutDurationGrid {
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 2rem;
    }

    input::-webkit-outer-spin-button,
    input::-webkit-inner-spin-button {
        -webkit-appearance: none;
        margin: 0;
    }
</style>
