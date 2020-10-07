/**
 * @fileoverview gRPC-Web generated client stub for v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.v1 = require('./entexaminfo_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.v1.EntExamInfoServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.v1.EntExamInfoServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.v1.GetEntExamInfoRequest,
 *   !proto.v1.GetEntExamInfoResponse>}
 */
const methodDescriptor_EntExamInfoService_GetEntExamInfo = new grpc.web.MethodDescriptor(
  '/v1.EntExamInfoService/GetEntExamInfo',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.v1.GetEntExamInfoRequest,
  proto.v1.GetEntExamInfoResponse,
  /**
   * @param {!proto.v1.GetEntExamInfoRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.v1.GetEntExamInfoResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.v1.GetEntExamInfoRequest,
 *   !proto.v1.GetEntExamInfoResponse>}
 */
const methodInfo_EntExamInfoService_GetEntExamInfo = new grpc.web.AbstractClientBase.MethodInfo(
  proto.v1.GetEntExamInfoResponse,
  /**
   * @param {!proto.v1.GetEntExamInfoRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.v1.GetEntExamInfoResponse.deserializeBinary
);


/**
 * @param {!proto.v1.GetEntExamInfoRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.v1.GetEntExamInfoResponse>}
 *     The XHR Node Readable Stream
 */
proto.v1.EntExamInfoServiceClient.prototype.getEntExamInfo =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/v1.EntExamInfoService/GetEntExamInfo',
      request,
      metadata || {},
      methodDescriptor_EntExamInfoService_GetEntExamInfo);
};


/**
 * @param {!proto.v1.GetEntExamInfoRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.v1.GetEntExamInfoResponse>}
 *     The XHR Node Readable Stream
 */
proto.v1.EntExamInfoServicePromiseClient.prototype.getEntExamInfo =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/v1.EntExamInfoService/GetEntExamInfo',
      request,
      metadata || {},
      methodDescriptor_EntExamInfoService_GetEntExamInfo);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.v1.GetEntExamPartRequest,
 *   !proto.v1.GetEntExamPartResponse>}
 */
const methodDescriptor_EntExamInfoService_GetEntExamPart = new grpc.web.MethodDescriptor(
  '/v1.EntExamInfoService/GetEntExamPart',
  grpc.web.MethodType.UNARY,
  proto.v1.GetEntExamPartRequest,
  proto.v1.GetEntExamPartResponse,
  /**
   * @param {!proto.v1.GetEntExamPartRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.v1.GetEntExamPartResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.v1.GetEntExamPartRequest,
 *   !proto.v1.GetEntExamPartResponse>}
 */
const methodInfo_EntExamInfoService_GetEntExamPart = new grpc.web.AbstractClientBase.MethodInfo(
  proto.v1.GetEntExamPartResponse,
  /**
   * @param {!proto.v1.GetEntExamPartRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.v1.GetEntExamPartResponse.deserializeBinary
);


/**
 * @param {!proto.v1.GetEntExamPartRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.v1.GetEntExamPartResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.v1.GetEntExamPartResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.v1.EntExamInfoServiceClient.prototype.getEntExamPart =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/v1.EntExamInfoService/GetEntExamPart',
      request,
      metadata || {},
      methodDescriptor_EntExamInfoService_GetEntExamPart,
      callback);
};


/**
 * @param {!proto.v1.GetEntExamPartRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.v1.GetEntExamPartResponse>}
 *     A native promise that resolves to the response
 */
proto.v1.EntExamInfoServicePromiseClient.prototype.getEntExamPart =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/v1.EntExamInfoService/GetEntExamPart',
      request,
      metadata || {},
      methodDescriptor_EntExamInfoService_GetEntExamPart);
};


module.exports = proto.v1;

