// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: blc.proto

package com.topview.api;

/**
 * Protobuf type {@code com.topview.api.GetDcFromActivityDTO}
 */
public final class GetDcFromActivityDTO extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:com.topview.api.GetDcFromActivityDTO)
    GetDcFromActivityDTOOrBuilder {
private static final long serialVersionUID = 0L;
  // Use GetDcFromActivityDTO.newBuilder() to construct.
  private GetDcFromActivityDTO(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private GetDcFromActivityDTO() {
    password_ = com.google.protobuf.ByteString.EMPTY;
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new GetDcFromActivityDTO();
  }

  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return com.topview.api.BlcServiceProto.internal_static_com_topview_api_GetDcFromActivityDTO_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return com.topview.api.BlcServiceProto.internal_static_com_topview_api_GetDcFromActivityDTO_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            com.topview.api.GetDcFromActivityDTO.class, com.topview.api.GetDcFromActivityDTO.Builder.class);
  }

  public static final int ACTIVITY_ID_FIELD_NUMBER = 1;
  private long activityId_ = 0L;
  /**
   * <code>int64 activity_id = 1;</code>
   * @return The activityId.
   */
  @java.lang.Override
  public long getActivityId() {
    return activityId_;
  }

  public static final int PASSWORD_FIELD_NUMBER = 2;
  private com.google.protobuf.ByteString password_ = com.google.protobuf.ByteString.EMPTY;
  /**
   * <code>bytes password = 2;</code>
   * @return The password.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString getPassword() {
    return password_;
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (activityId_ != 0L) {
      output.writeInt64(1, activityId_);
    }
    if (!password_.isEmpty()) {
      output.writeBytes(2, password_);
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (activityId_ != 0L) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt64Size(1, activityId_);
    }
    if (!password_.isEmpty()) {
      size += com.google.protobuf.CodedOutputStream
        .computeBytesSize(2, password_);
    }
    size += getUnknownFields().getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof com.topview.api.GetDcFromActivityDTO)) {
      return super.equals(obj);
    }
    com.topview.api.GetDcFromActivityDTO other = (com.topview.api.GetDcFromActivityDTO) obj;

    if (getActivityId()
        != other.getActivityId()) return false;
    if (!getPassword()
        .equals(other.getPassword())) return false;
    if (!getUnknownFields().equals(other.getUnknownFields())) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + ACTIVITY_ID_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashLong(
        getActivityId());
    hash = (37 * hash) + PASSWORD_FIELD_NUMBER;
    hash = (53 * hash) + getPassword().hashCode();
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static com.topview.api.GetDcFromActivityDTO parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.topview.api.GetDcFromActivityDTO parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.topview.api.GetDcFromActivityDTO parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.topview.api.GetDcFromActivityDTO parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.topview.api.GetDcFromActivityDTO parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.topview.api.GetDcFromActivityDTO parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.topview.api.GetDcFromActivityDTO parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.topview.api.GetDcFromActivityDTO parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.topview.api.GetDcFromActivityDTO parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static com.topview.api.GetDcFromActivityDTO parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.topview.api.GetDcFromActivityDTO parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.topview.api.GetDcFromActivityDTO parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(com.topview.api.GetDcFromActivityDTO prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code com.topview.api.GetDcFromActivityDTO}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:com.topview.api.GetDcFromActivityDTO)
      com.topview.api.GetDcFromActivityDTOOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.topview.api.BlcServiceProto.internal_static_com_topview_api_GetDcFromActivityDTO_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.topview.api.BlcServiceProto.internal_static_com_topview_api_GetDcFromActivityDTO_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.topview.api.GetDcFromActivityDTO.class, com.topview.api.GetDcFromActivityDTO.Builder.class);
    }

    // Construct using com.topview.api.GetDcFromActivityDTO.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      bitField0_ = 0;
      activityId_ = 0L;
      password_ = com.google.protobuf.ByteString.EMPTY;
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return com.topview.api.BlcServiceProto.internal_static_com_topview_api_GetDcFromActivityDTO_descriptor;
    }

    @java.lang.Override
    public com.topview.api.GetDcFromActivityDTO getDefaultInstanceForType() {
      return com.topview.api.GetDcFromActivityDTO.getDefaultInstance();
    }

    @java.lang.Override
    public com.topview.api.GetDcFromActivityDTO build() {
      com.topview.api.GetDcFromActivityDTO result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public com.topview.api.GetDcFromActivityDTO buildPartial() {
      com.topview.api.GetDcFromActivityDTO result = new com.topview.api.GetDcFromActivityDTO(this);
      if (bitField0_ != 0) { buildPartial0(result); }
      onBuilt();
      return result;
    }

    private void buildPartial0(com.topview.api.GetDcFromActivityDTO result) {
      int from_bitField0_ = bitField0_;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        result.activityId_ = activityId_;
      }
      if (((from_bitField0_ & 0x00000002) != 0)) {
        result.password_ = password_;
      }
    }

    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof com.topview.api.GetDcFromActivityDTO) {
        return mergeFrom((com.topview.api.GetDcFromActivityDTO)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(com.topview.api.GetDcFromActivityDTO other) {
      if (other == com.topview.api.GetDcFromActivityDTO.getDefaultInstance()) return this;
      if (other.getActivityId() != 0L) {
        setActivityId(other.getActivityId());
      }
      if (other.getPassword() != com.google.protobuf.ByteString.EMPTY) {
        setPassword(other.getPassword());
      }
      this.mergeUnknownFields(other.getUnknownFields());
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 8: {
              activityId_ = input.readInt64();
              bitField0_ |= 0x00000001;
              break;
            } // case 8
            case 18: {
              password_ = input.readBytes();
              bitField0_ |= 0x00000002;
              break;
            } // case 18
            default: {
              if (!super.parseUnknownField(input, extensionRegistry, tag)) {
                done = true; // was an endgroup tag
              }
              break;
            } // default:
          } // switch (tag)
        } // while (!done)
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.unwrapIOException();
      } finally {
        onChanged();
      } // finally
      return this;
    }
    private int bitField0_;

    private long activityId_ ;
    /**
     * <code>int64 activity_id = 1;</code>
     * @return The activityId.
     */
    @java.lang.Override
    public long getActivityId() {
      return activityId_;
    }
    /**
     * <code>int64 activity_id = 1;</code>
     * @param value The activityId to set.
     * @return This builder for chaining.
     */
    public Builder setActivityId(long value) {

      activityId_ = value;
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <code>int64 activity_id = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearActivityId() {
      bitField0_ = (bitField0_ & ~0x00000001);
      activityId_ = 0L;
      onChanged();
      return this;
    }

    private com.google.protobuf.ByteString password_ = com.google.protobuf.ByteString.EMPTY;
    /**
     * <code>bytes password = 2;</code>
     * @return The password.
     */
    @java.lang.Override
    public com.google.protobuf.ByteString getPassword() {
      return password_;
    }
    /**
     * <code>bytes password = 2;</code>
     * @param value The password to set.
     * @return This builder for chaining.
     */
    public Builder setPassword(com.google.protobuf.ByteString value) {
      if (value == null) { throw new NullPointerException(); }
      password_ = value;
      bitField0_ |= 0x00000002;
      onChanged();
      return this;
    }
    /**
     * <code>bytes password = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearPassword() {
      bitField0_ = (bitField0_ & ~0x00000002);
      password_ = getDefaultInstance().getPassword();
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:com.topview.api.GetDcFromActivityDTO)
  }

  // @@protoc_insertion_point(class_scope:com.topview.api.GetDcFromActivityDTO)
  private static final com.topview.api.GetDcFromActivityDTO DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new com.topview.api.GetDcFromActivityDTO();
  }

  public static com.topview.api.GetDcFromActivityDTO getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<GetDcFromActivityDTO>
      PARSER = new com.google.protobuf.AbstractParser<GetDcFromActivityDTO>() {
    @java.lang.Override
    public GetDcFromActivityDTO parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      Builder builder = newBuilder();
      try {
        builder.mergeFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(builder.buildPartial());
      } catch (com.google.protobuf.UninitializedMessageException e) {
        throw e.asInvalidProtocolBufferException().setUnfinishedMessage(builder.buildPartial());
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(e)
            .setUnfinishedMessage(builder.buildPartial());
      }
      return builder.buildPartial();
    }
  };

  public static com.google.protobuf.Parser<GetDcFromActivityDTO> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<GetDcFromActivityDTO> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public com.topview.api.GetDcFromActivityDTO getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

