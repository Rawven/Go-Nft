// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: blc.proto

package com.topview.api;

/**
 * Protobuf type {@code com.topview.api.CheckDcAndReturnTimeRequest}
 */
public final class CheckDcAndReturnTimeRequest extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:com.topview.api.CheckDcAndReturnTimeRequest)
    CheckDcAndReturnTimeRequestOrBuilder {
private static final long serialVersionUID = 0L;
  // Use CheckDcAndReturnTimeRequest.newBuilder() to construct.
  private CheckDcAndReturnTimeRequest(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private CheckDcAndReturnTimeRequest() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new CheckDcAndReturnTimeRequest();
  }

  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return com.topview.api.BlcServiceProto.internal_static_com_topview_api_CheckDcAndReturnTimeRequest_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return com.topview.api.BlcServiceProto.internal_static_com_topview_api_CheckDcAndReturnTimeRequest_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            com.topview.api.CheckDcAndReturnTimeRequest.class, com.topview.api.CheckDcAndReturnTimeRequest.Builder.class);
  }

  public static final int DTO_FIELD_NUMBER = 1;
  private com.topview.api.CheckDcAndReturnTimeDTO dto_;
  /**
   * <code>.com.topview.api.CheckDcAndReturnTimeDTO dto = 1;</code>
   * @return Whether the dto field is set.
   */
  @java.lang.Override
  public boolean hasDto() {
    return dto_ != null;
  }
  /**
   * <code>.com.topview.api.CheckDcAndReturnTimeDTO dto = 1;</code>
   * @return The dto.
   */
  @java.lang.Override
  public com.topview.api.CheckDcAndReturnTimeDTO getDto() {
    return dto_ == null ? com.topview.api.CheckDcAndReturnTimeDTO.getDefaultInstance() : dto_;
  }
  /**
   * <code>.com.topview.api.CheckDcAndReturnTimeDTO dto = 1;</code>
   */
  @java.lang.Override
  public com.topview.api.CheckDcAndReturnTimeDTOOrBuilder getDtoOrBuilder() {
    return dto_ == null ? com.topview.api.CheckDcAndReturnTimeDTO.getDefaultInstance() : dto_;
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
    if (dto_ != null) {
      output.writeMessage(1, getDto());
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (dto_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(1, getDto());
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
    if (!(obj instanceof com.topview.api.CheckDcAndReturnTimeRequest)) {
      return super.equals(obj);
    }
    com.topview.api.CheckDcAndReturnTimeRequest other = (com.topview.api.CheckDcAndReturnTimeRequest) obj;

    if (hasDto() != other.hasDto()) return false;
    if (hasDto()) {
      if (!getDto()
          .equals(other.getDto())) return false;
    }
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
    if (hasDto()) {
      hash = (37 * hash) + DTO_FIELD_NUMBER;
      hash = (53 * hash) + getDto().hashCode();
    }
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static com.topview.api.CheckDcAndReturnTimeRequest parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.topview.api.CheckDcAndReturnTimeRequest parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.topview.api.CheckDcAndReturnTimeRequest parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.topview.api.CheckDcAndReturnTimeRequest parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.topview.api.CheckDcAndReturnTimeRequest parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.topview.api.CheckDcAndReturnTimeRequest parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.topview.api.CheckDcAndReturnTimeRequest parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.topview.api.CheckDcAndReturnTimeRequest parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.topview.api.CheckDcAndReturnTimeRequest parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static com.topview.api.CheckDcAndReturnTimeRequest parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.topview.api.CheckDcAndReturnTimeRequest parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.topview.api.CheckDcAndReturnTimeRequest parseFrom(
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
  public static Builder newBuilder(com.topview.api.CheckDcAndReturnTimeRequest prototype) {
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
   * Protobuf type {@code com.topview.api.CheckDcAndReturnTimeRequest}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:com.topview.api.CheckDcAndReturnTimeRequest)
      com.topview.api.CheckDcAndReturnTimeRequestOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.topview.api.BlcServiceProto.internal_static_com_topview_api_CheckDcAndReturnTimeRequest_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.topview.api.BlcServiceProto.internal_static_com_topview_api_CheckDcAndReturnTimeRequest_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.topview.api.CheckDcAndReturnTimeRequest.class, com.topview.api.CheckDcAndReturnTimeRequest.Builder.class);
    }

    // Construct using com.topview.api.CheckDcAndReturnTimeRequest.newBuilder()
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
      dto_ = null;
      if (dtoBuilder_ != null) {
        dtoBuilder_.dispose();
        dtoBuilder_ = null;
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return com.topview.api.BlcServiceProto.internal_static_com_topview_api_CheckDcAndReturnTimeRequest_descriptor;
    }

    @java.lang.Override
    public com.topview.api.CheckDcAndReturnTimeRequest getDefaultInstanceForType() {
      return com.topview.api.CheckDcAndReturnTimeRequest.getDefaultInstance();
    }

    @java.lang.Override
    public com.topview.api.CheckDcAndReturnTimeRequest build() {
      com.topview.api.CheckDcAndReturnTimeRequest result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public com.topview.api.CheckDcAndReturnTimeRequest buildPartial() {
      com.topview.api.CheckDcAndReturnTimeRequest result = new com.topview.api.CheckDcAndReturnTimeRequest(this);
      if (bitField0_ != 0) { buildPartial0(result); }
      onBuilt();
      return result;
    }

    private void buildPartial0(com.topview.api.CheckDcAndReturnTimeRequest result) {
      int from_bitField0_ = bitField0_;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        result.dto_ = dtoBuilder_ == null
            ? dto_
            : dtoBuilder_.build();
      }
    }

    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof com.topview.api.CheckDcAndReturnTimeRequest) {
        return mergeFrom((com.topview.api.CheckDcAndReturnTimeRequest)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(com.topview.api.CheckDcAndReturnTimeRequest other) {
      if (other == com.topview.api.CheckDcAndReturnTimeRequest.getDefaultInstance()) return this;
      if (other.hasDto()) {
        mergeDto(other.getDto());
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
            case 10: {
              input.readMessage(
                  getDtoFieldBuilder().getBuilder(),
                  extensionRegistry);
              bitField0_ |= 0x00000001;
              break;
            } // case 10
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

    private com.topview.api.CheckDcAndReturnTimeDTO dto_;
    private com.google.protobuf.SingleFieldBuilderV3<
        com.topview.api.CheckDcAndReturnTimeDTO, com.topview.api.CheckDcAndReturnTimeDTO.Builder, com.topview.api.CheckDcAndReturnTimeDTOOrBuilder> dtoBuilder_;
    /**
     * <code>.com.topview.api.CheckDcAndReturnTimeDTO dto = 1;</code>
     * @return Whether the dto field is set.
     */
    public boolean hasDto() {
      return ((bitField0_ & 0x00000001) != 0);
    }
    /**
     * <code>.com.topview.api.CheckDcAndReturnTimeDTO dto = 1;</code>
     * @return The dto.
     */
    public com.topview.api.CheckDcAndReturnTimeDTO getDto() {
      if (dtoBuilder_ == null) {
        return dto_ == null ? com.topview.api.CheckDcAndReturnTimeDTO.getDefaultInstance() : dto_;
      } else {
        return dtoBuilder_.getMessage();
      }
    }
    /**
     * <code>.com.topview.api.CheckDcAndReturnTimeDTO dto = 1;</code>
     */
    public Builder setDto(com.topview.api.CheckDcAndReturnTimeDTO value) {
      if (dtoBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        dto_ = value;
      } else {
        dtoBuilder_.setMessage(value);
      }
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <code>.com.topview.api.CheckDcAndReturnTimeDTO dto = 1;</code>
     */
    public Builder setDto(
        com.topview.api.CheckDcAndReturnTimeDTO.Builder builderForValue) {
      if (dtoBuilder_ == null) {
        dto_ = builderForValue.build();
      } else {
        dtoBuilder_.setMessage(builderForValue.build());
      }
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <code>.com.topview.api.CheckDcAndReturnTimeDTO dto = 1;</code>
     */
    public Builder mergeDto(com.topview.api.CheckDcAndReturnTimeDTO value) {
      if (dtoBuilder_ == null) {
        if (((bitField0_ & 0x00000001) != 0) &&
          dto_ != null &&
          dto_ != com.topview.api.CheckDcAndReturnTimeDTO.getDefaultInstance()) {
          getDtoBuilder().mergeFrom(value);
        } else {
          dto_ = value;
        }
      } else {
        dtoBuilder_.mergeFrom(value);
      }
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <code>.com.topview.api.CheckDcAndReturnTimeDTO dto = 1;</code>
     */
    public Builder clearDto() {
      bitField0_ = (bitField0_ & ~0x00000001);
      dto_ = null;
      if (dtoBuilder_ != null) {
        dtoBuilder_.dispose();
        dtoBuilder_ = null;
      }
      onChanged();
      return this;
    }
    /**
     * <code>.com.topview.api.CheckDcAndReturnTimeDTO dto = 1;</code>
     */
    public com.topview.api.CheckDcAndReturnTimeDTO.Builder getDtoBuilder() {
      bitField0_ |= 0x00000001;
      onChanged();
      return getDtoFieldBuilder().getBuilder();
    }
    /**
     * <code>.com.topview.api.CheckDcAndReturnTimeDTO dto = 1;</code>
     */
    public com.topview.api.CheckDcAndReturnTimeDTOOrBuilder getDtoOrBuilder() {
      if (dtoBuilder_ != null) {
        return dtoBuilder_.getMessageOrBuilder();
      } else {
        return dto_ == null ?
            com.topview.api.CheckDcAndReturnTimeDTO.getDefaultInstance() : dto_;
      }
    }
    /**
     * <code>.com.topview.api.CheckDcAndReturnTimeDTO dto = 1;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        com.topview.api.CheckDcAndReturnTimeDTO, com.topview.api.CheckDcAndReturnTimeDTO.Builder, com.topview.api.CheckDcAndReturnTimeDTOOrBuilder> 
        getDtoFieldBuilder() {
      if (dtoBuilder_ == null) {
        dtoBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            com.topview.api.CheckDcAndReturnTimeDTO, com.topview.api.CheckDcAndReturnTimeDTO.Builder, com.topview.api.CheckDcAndReturnTimeDTOOrBuilder>(
                getDto(),
                getParentForChildren(),
                isClean());
        dto_ = null;
      }
      return dtoBuilder_;
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


    // @@protoc_insertion_point(builder_scope:com.topview.api.CheckDcAndReturnTimeRequest)
  }

  // @@protoc_insertion_point(class_scope:com.topview.api.CheckDcAndReturnTimeRequest)
  private static final com.topview.api.CheckDcAndReturnTimeRequest DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new com.topview.api.CheckDcAndReturnTimeRequest();
  }

  public static com.topview.api.CheckDcAndReturnTimeRequest getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<CheckDcAndReturnTimeRequest>
      PARSER = new com.google.protobuf.AbstractParser<CheckDcAndReturnTimeRequest>() {
    @java.lang.Override
    public CheckDcAndReturnTimeRequest parsePartialFrom(
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

  public static com.google.protobuf.Parser<CheckDcAndReturnTimeRequest> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<CheckDcAndReturnTimeRequest> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public com.topview.api.CheckDcAndReturnTimeRequest getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

