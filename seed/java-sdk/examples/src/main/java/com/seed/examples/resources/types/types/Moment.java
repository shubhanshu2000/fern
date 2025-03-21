/**
 * This file was auto-generated by Fern from our API Definition.
 */
package com.seed.examples.resources.types.types;

import com.fasterxml.jackson.annotation.JsonAnyGetter;
import com.fasterxml.jackson.annotation.JsonAnySetter;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.seed.examples.core.ObjectMappers;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
import java.util.UUID;
import org.jetbrains.annotations.NotNull;

@JsonInclude(JsonInclude.Include.NON_ABSENT)
@JsonDeserialize(builder = Moment.Builder.class)
public final class Moment {
    private final UUID id;

    private final String date;

    private final OffsetDateTime datetime;

    private final Map<String, Object> additionalProperties;

    private Moment(UUID id, String date, OffsetDateTime datetime, Map<String, Object> additionalProperties) {
        this.id = id;
        this.date = date;
        this.datetime = datetime;
        this.additionalProperties = additionalProperties;
    }

    @JsonProperty("id")
    public UUID getId() {
        return id;
    }

    @JsonProperty("date")
    public String getDate() {
        return date;
    }

    @JsonProperty("datetime")
    public OffsetDateTime getDatetime() {
        return datetime;
    }

    @java.lang.Override
    public boolean equals(Object other) {
        if (this == other) return true;
        return other instanceof Moment && equalTo((Moment) other);
    }

    @JsonAnyGetter
    public Map<String, Object> getAdditionalProperties() {
        return this.additionalProperties;
    }

    private boolean equalTo(Moment other) {
        return id.equals(other.id) && date.equals(other.date) && datetime.equals(other.datetime);
    }

    @java.lang.Override
    public int hashCode() {
        return Objects.hash(this.id, this.date, this.datetime);
    }

    @java.lang.Override
    public String toString() {
        return ObjectMappers.stringify(this);
    }

    public static IdStage builder() {
        return new Builder();
    }

    public interface IdStage {
        DateStage id(@NotNull UUID id);

        Builder from(Moment other);
    }

    public interface DateStage {
        DatetimeStage date(@NotNull String date);
    }

    public interface DatetimeStage {
        _FinalStage datetime(@NotNull OffsetDateTime datetime);
    }

    public interface _FinalStage {
        Moment build();
    }

    @JsonIgnoreProperties(ignoreUnknown = true)
    public static final class Builder implements IdStage, DateStage, DatetimeStage, _FinalStage {
        private UUID id;

        private String date;

        private OffsetDateTime datetime;

        @JsonAnySetter
        private Map<String, Object> additionalProperties = new HashMap<>();

        private Builder() {}

        @java.lang.Override
        public Builder from(Moment other) {
            id(other.getId());
            date(other.getDate());
            datetime(other.getDatetime());
            return this;
        }

        @java.lang.Override
        @JsonSetter("id")
        public DateStage id(@NotNull UUID id) {
            this.id = Objects.requireNonNull(id, "id must not be null");
            return this;
        }

        @java.lang.Override
        @JsonSetter("date")
        public DatetimeStage date(@NotNull String date) {
            this.date = Objects.requireNonNull(date, "date must not be null");
            return this;
        }

        @java.lang.Override
        @JsonSetter("datetime")
        public _FinalStage datetime(@NotNull OffsetDateTime datetime) {
            this.datetime = Objects.requireNonNull(datetime, "datetime must not be null");
            return this;
        }

        @java.lang.Override
        public Moment build() {
            return new Moment(id, date, datetime, additionalProperties);
        }
    }
}
