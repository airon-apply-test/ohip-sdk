/**
 * OPERA Cloud Xchange Interface OXI API
 * APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 23.0.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import * as runtime from '../runtime';
import type { ChangedAutomaticTransmissionSchedules, ChangedCommunicationMethods, ChangedConversionCodeMappings, ChangedConversionCodesStatus, ChangedExternalInterfaceSetups, ChangedInterfaceControls, ChangedProfileMatchRules, CreatedConversionCodeMappings, FetchAccumulatedBusinessEvents, FetchAutomaticTransmissionSchedules, FetchCommunicationMethods, FetchConversionCodeMappings, FetchConversionCodes, FetchExternalInterfaceSetups, FetchIntegrationInboundMessages, FetchIntegrationOutboundMessages, FetchInterfaceControls, FetchOXIListOfValues, FetchProfileMatchRules, LegacyInterfaceStatusDetails, PostAutomaticTransmissionSchedulesRequest, PostConversionCodeMappingsRequest, PostExternalInterfaceSetupsRequest, PostProfileMatchRulesRequest, PutAutomaticTransmissionSchedulesRequest, PutCommunicationMethodsRequest, PutConversionCodeMappingsRequest, PutConversionCodesStatusRequest, PutExternalInterfaceSetupsRequest, PutInterfaceControlsRequest, PutProfileMatchRulesRequest, Status } from '../models';
export interface DeleteAutomaticTransmissionSchedulesRequest {
    messageId?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface DeleteConversionCodeMappingsRequest {
    id?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface DeleteExternalInterfaceSetupsRequest {
    interfaceId?: string;
    hotelId?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface DeleteProfileMatchRulesRequest {
    profileType?: string;
    interfaceId?: string;
    hotelId?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface DeleteUDFMappingsRequest {
    operaValue?: string;
    conversionCode?: string;
    interfaceId?: string;
    hotelId?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    profileType?: string;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface GetAccumulatedBusinessEventsRequest {
    interfaceId?: string;
    hotelId?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    limit?: number;
    offset?: number;
    modules?: Set<GetAccumulatedBusinessEventsModulesEnum>;
    endDate?: Date;
    startDate?: Date;
    primaryKey?: string;
    status?: GetAccumulatedBusinessEventsStatusEnum;
    hotelIds?: Array<string>;
    integrationSystem?: GetAccumulatedBusinessEventsIntegrationSystemEnum;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface GetAutomaticTransmissionSchedulesRequest {
    hotelId?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    hotelIds?: Array<string>;
    interfaceList?: Array<string>;
    includeInactive?: boolean;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface GetCommunicationMethodsRequest {
    interfaceId?: string;
    isGlobal?: boolean;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    hotelId?: Array<string>;
    interfaceIds?: Array<string>;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface GetConversionCodeMappingsRequest {
    conversionCode?: string;
    interfaceId?: string;
    hotelId?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    integrationSystem?: GetConversionCodeMappingsIntegrationSystemEnum;
    udfCode?: boolean;
    mappingValue?: string;
    searchBy?: GetConversionCodeMappingsSearchByEnum;
    includeGlobal?: boolean;
    profileType?: string;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface GetConversionCodesRequest {
    interfaceId?: string;
    hotelId?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    integrationSystem?: GetConversionCodesIntegrationSystemEnum;
    conversionCode?: string;
    udfCode?: boolean;
    includeInactive?: boolean;
    group?: string;
    includeUDFCodes?: boolean;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface GetExternalInterfaceSetupsRequest {
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    interfaceIds?: Array<string>;
    hotelIds?: Array<string>;
    includeXmlVersion?: boolean;
    systemType?: GetExternalInterfaceSetupsSystemTypeEnum;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface GetIntegrationInboundMessagesRequest {
    integrationSystem?: string;
    interfaceId?: string;
    externalHotelCode?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    externalHotelCodes?: Array<string>;
    messageTypes?: Array<string>;
    endDate?: Date;
    startDate?: Date;
    includeReviewed?: boolean;
    fromMessageID?: string;
    toMessageID?: string;
    messageStatus?: Array<string>;
    messageReference?: string;
    errorMessageWildCard?: string;
    errorMessageType?: GetIntegrationInboundMessagesErrorMessageTypeEnum;
    msgContains?: string;
    anyOneOfTheMsg?: boolean;
    valuesOnly?: boolean;
    limit?: number;
    offset?: number;
    includeErrors?: boolean;
    actionId?: string;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface GetIntegrationOutboundMessagesRequest {
    integrationSystem?: string;
    interfaceId?: string;
    hotelId?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    hotelIds?: Array<string>;
    actionType?: string;
    errorMessageWildCard?: string;
    errorMessageType?: GetIntegrationOutboundMessagesErrorMessageTypeEnum;
    endDate?: Date;
    startDate?: Date;
    includeReviewed?: boolean;
    fromMessageID?: string;
    toMessageID?: string;
    messageStatus?: Array<string>;
    msgContains?: string;
    anyOneOfTheMsg?: boolean;
    valuesOnly?: boolean;
    messageReference?: string;
    messageTypes?: Array<string>;
    limit?: number;
    offset?: number;
    includeErrors?: boolean;
    actionId?: string;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface GetInterfaceControlsRequest {
    interfaceId?: string;
    hotelId?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    integrationSystem?: GetInterfaceControlsIntegrationSystemEnum;
    croCode?: string;
    parameterNameWildCard?: string;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface GetLegacyInterfaceStatusRequest {
    interfaceId?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface GetOXIListOfValuesRequest {
    lovCode?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    integrationSystem?: GetOXIListOfValuesIntegrationSystemEnum;
    includeInActive?: GetOXIListOfValuesIncludeInActiveEnum;
    parameterNames?: Array<string>;
    parameterValues?: Array<string>;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface GetProfileMatchRulesRequest {
    interfaceId?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    hotelId?: Array<string>;
    includeXmlVersions?: boolean;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface PostAutomaticTransmissionSchedulesOperationRequest {
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    automaticTransmissionSchedulesToBeCreated?: PostAutomaticTransmissionSchedulesRequest;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface PostConversionCodeMappingsOperationRequest {
    externalValue?: string;
    operaValue?: string;
    conversionCode?: string;
    interfaceId?: string;
    hotelId?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    conversionCodeMappingsToBeCreated?: PostConversionCodeMappingsRequest;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface PostExternalInterfaceSetupsOperationRequest {
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    externalInterfaceSetupsToBeCreated?: PostExternalInterfaceSetupsRequest;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface PostProfileMatchRulesOperationRequest {
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    profileMatchRulesToBeCreated?: PostProfileMatchRulesRequest;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface PutAutomaticTransmissionSchedulesOperationRequest {
    hotelId?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    automaticTransmissionSchedulesToBeChanged?: PutAutomaticTransmissionSchedulesRequest;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface PutCommunicationMethodsOperationRequest {
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    communicationMethodsToBeChanged?: PutCommunicationMethodsRequest;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface PutConversionCodeMappingsOperationRequest {
    externalValue?: string;
    operaValue?: string;
    id?: string;
    conversionCode?: string;
    interfaceId?: string;
    hotelId?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    conversionCodeMappingsToBeChanged?: PutConversionCodeMappingsRequest;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface PutConversionCodesStatusOperationRequest {
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    conversionCodesStatusToBeChanged?: PutConversionCodesStatusRequest;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface PutExternalInterfaceSetupsOperationRequest {
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    externalInterfaceSetupsToBeChanged?: PutExternalInterfaceSetupsRequest;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface PutInterfaceControlsOperationRequest {
    interfaceId?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    interfaceControlsToBeChanged?: PutInterfaceControlsRequest;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface PutProfileMatchRulesOperationRequest {
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    profileMatchRulesToBeChanged?: PutProfileMatchRulesRequest;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
/**
 *
 */
export declare class OperaExchangeInterfaceConfigApi extends runtime.BaseAPI {
    /**
     * API to Delete  Automatic Transmission Schedules by Scheduled Message Id. <p><strong>OperationId:</strong>deleteAutomaticTransmissionSchedules</p>
     * Remove Automatic Transmission Schedules
     */
    deleteAutomaticTransmissionSchedulesRaw(requestParameters: DeleteAutomaticTransmissionSchedulesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Status>>;
    /**
     * API to Delete  Automatic Transmission Schedules by Scheduled Message Id. <p><strong>OperationId:</strong>deleteAutomaticTransmissionSchedules</p>
     * Remove Automatic Transmission Schedules
     */
    deleteAutomaticTransmissionSchedules(requestParameters: DeleteAutomaticTransmissionSchedulesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Status>;
    /**
     * API to Delete  Interface Mapping Conversion Codes by Conversion Code Mapping Id. <p><strong>OperationId:</strong>deleteConversionCodeMappings</p>
     * Remove Conversion Codes Mappings
     */
    deleteConversionCodeMappingsRaw(requestParameters: DeleteConversionCodeMappingsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Status>>;
    /**
     * API to Delete  Interface Mapping Conversion Codes by Conversion Code Mapping Id. <p><strong>OperationId:</strong>deleteConversionCodeMappings</p>
     * Remove Conversion Codes Mappings
     */
    deleteConversionCodeMappings(requestParameters: DeleteConversionCodeMappingsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Status>;
    /**
     * API to Delete Interface Setup configuration by Hotel Id and Interface Id. <p><strong>OperationId:</strong>deleteExternalInterfaceSetups</p>
     * Delete external Interface Setup details
     */
    deleteExternalInterfaceSetupsRaw(requestParameters: DeleteExternalInterfaceSetupsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Status>>;
    /**
     * API to Delete Interface Setup configuration by Hotel Id and Interface Id. <p><strong>OperationId:</strong>deleteExternalInterfaceSetups</p>
     * Delete external Interface Setup details
     */
    deleteExternalInterfaceSetups(requestParameters: DeleteExternalInterfaceSetupsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Status>;
    /**
     * API to Delete  Profile Match Rules by Hotel Id, Interface Id and Profile Type. <p><strong>OperationId:</strong>deleteProfileMatchRules</p>
     * Delete profile match rules
     */
    deleteProfileMatchRulesRaw(requestParameters: DeleteProfileMatchRulesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Status>>;
    /**
     * API to Delete  Profile Match Rules by Hotel Id, Interface Id and Profile Type. <p><strong>OperationId:</strong>deleteProfileMatchRules</p>
     * Delete profile match rules
     */
    deleteProfileMatchRules(requestParameters: DeleteProfileMatchRulesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Status>;
    /**
     * API to Delete Interface Mapping UDF Conversion Codes by Hotel Id, Interface Id and Conversion Code. <p><strong>OperationId:</strong>deleteUDFMappings</p>
     * Remove UDF Mappings
     */
    deleteUDFMappingsRaw(requestParameters: DeleteUDFMappingsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Status>>;
    /**
     * API to Delete Interface Mapping UDF Conversion Codes by Hotel Id, Interface Id and Conversion Code. <p><strong>OperationId:</strong>deleteUDFMappings</p>
     * Remove UDF Mappings
     */
    deleteUDFMappings(requestParameters: DeleteUDFMappingsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Status>;
    /**
     * API to Fetch  Accumulated Business Event Messages for External Systems by Hotel Id, Interface Id. <p><strong>OperationId:</strong>getAccumulatedBusinessEvents</p>
     * Fetch accumulated business events
     */
    getAccumulatedBusinessEventsRaw(requestParameters: GetAccumulatedBusinessEventsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<FetchAccumulatedBusinessEvents>>;
    /**
     * API to Fetch  Accumulated Business Event Messages for External Systems by Hotel Id, Interface Id. <p><strong>OperationId:</strong>getAccumulatedBusinessEvents</p>
     * Fetch accumulated business events
     */
    getAccumulatedBusinessEvents(requestParameters: GetAccumulatedBusinessEventsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<FetchAccumulatedBusinessEvents>;
    /**
     * API to Fetch Automatic Transmission Schedules by Hotel Id. <p><strong>OperationId:</strong>getAutomaticTransmissionSchedules</p>
     * Fetch Automatic Transmission Schedules
     */
    getAutomaticTransmissionSchedulesRaw(requestParameters: GetAutomaticTransmissionSchedulesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<FetchAutomaticTransmissionSchedules>>;
    /**
     * API to Fetch Automatic Transmission Schedules by Hotel Id. <p><strong>OperationId:</strong>getAutomaticTransmissionSchedules</p>
     * Fetch Automatic Transmission Schedules
     */
    getAutomaticTransmissionSchedules(requestParameters: GetAutomaticTransmissionSchedulesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<FetchAutomaticTransmissionSchedules>;
    /**
     * API to Fetch Communication Methods. <p><strong>OperationId:</strong>getCommunicationMethods</p>
     * Fetch Communication Methods
     */
    getCommunicationMethodsRaw(requestParameters: GetCommunicationMethodsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<FetchCommunicationMethods>>;
    /**
     * API to Fetch Communication Methods. <p><strong>OperationId:</strong>getCommunicationMethods</p>
     * Fetch Communication Methods
     */
    getCommunicationMethods(requestParameters: GetCommunicationMethodsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<FetchCommunicationMethods>;
    /**
     * API to Fetch Interface Mapping  Conversion Codes by Hotel Id, Interface Id and Conversion Code. <p><strong>OperationId:</strong>getConversionCodeMappings</p>
     * Fetch Conversion Code Mappings
     */
    getConversionCodeMappingsRaw(requestParameters: GetConversionCodeMappingsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<FetchConversionCodeMappings>>;
    /**
     * API to Fetch Interface Mapping  Conversion Codes by Hotel Id, Interface Id and Conversion Code. <p><strong>OperationId:</strong>getConversionCodeMappings</p>
     * Fetch Conversion Code Mappings
     */
    getConversionCodeMappings(requestParameters: GetConversionCodeMappingsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<FetchConversionCodeMappings>;
    /**
     * API to Fetch Interface Mappings UDF Conversion Codes by Hotel Id and Interface Id. <p><strong>OperationId:</strong>getConversionCodes</p>
     * Fetch Conversion Codes
     */
    getConversionCodesRaw(requestParameters: GetConversionCodesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<FetchConversionCodes>>;
    /**
     * API to Fetch Interface Mappings UDF Conversion Codes by Hotel Id and Interface Id. <p><strong>OperationId:</strong>getConversionCodes</p>
     * Fetch Conversion Codes
     */
    getConversionCodes(requestParameters: GetConversionCodesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<FetchConversionCodes>;
    /**
     * API to Fetch Interface Setup configurations by Interface Ids and Hotel Ids. <p><strong>OperationId:</strong>getExternalInterfaceSetups</p>
     * fetch external Interface Setup details
     */
    getExternalInterfaceSetupsRaw(requestParameters: GetExternalInterfaceSetupsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<FetchExternalInterfaceSetups>>;
    /**
     * API to Fetch Interface Setup configurations by Interface Ids and Hotel Ids. <p><strong>OperationId:</strong>getExternalInterfaceSetups</p>
     * fetch external Interface Setup details
     */
    getExternalInterfaceSetups(requestParameters: GetExternalInterfaceSetupsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<FetchExternalInterfaceSetups>;
    /**
     * API to Fetch Inbound Messages From External System by Hotel Id, Interface Id. <p><strong>OperationId:</strong>getIntegrationInboundMessages</p>
     * Fetch integration inbound messages
     */
    getIntegrationInboundMessagesRaw(requestParameters: GetIntegrationInboundMessagesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<FetchIntegrationInboundMessages>>;
    /**
     * API to Fetch Inbound Messages From External System by Hotel Id, Interface Id. <p><strong>OperationId:</strong>getIntegrationInboundMessages</p>
     * Fetch integration inbound messages
     */
    getIntegrationInboundMessages(requestParameters: GetIntegrationInboundMessagesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<FetchIntegrationInboundMessages>;
    /**
     * API to Fetch Outbound Messages To External System by Hotel Id, Interface Id. <p><strong>OperationId:</strong>getIntegrationOutboundMessages</p>
     * fetch integration outbound messages
     */
    getIntegrationOutboundMessagesRaw(requestParameters: GetIntegrationOutboundMessagesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<FetchIntegrationOutboundMessages>>;
    /**
     * API to Fetch Outbound Messages To External System by Hotel Id, Interface Id. <p><strong>OperationId:</strong>getIntegrationOutboundMessages</p>
     * fetch integration outbound messages
     */
    getIntegrationOutboundMessages(requestParameters: GetIntegrationOutboundMessagesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<FetchIntegrationOutboundMessages>;
    /**
     * API to Fetch  Interface Controls  by Hotel Id and Interface Id <p><strong>OperationId:</strong>getInterfaceControls</p>
     * Retrieve OXI Parameters and Defaults
     */
    getInterfaceControlsRaw(requestParameters: GetInterfaceControlsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<FetchInterfaceControls>>;
    /**
     * API to Fetch  Interface Controls  by Hotel Id and Interface Id <p><strong>OperationId:</strong>getInterfaceControls</p>
     * Retrieve OXI Parameters and Defaults
     */
    getInterfaceControls(requestParameters: GetInterfaceControlsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<FetchInterfaceControls>;
    /**
     * API to Fetch legacy OXI Interface Processor Status by Interface Id. <p><strong>OperationId:</strong>getLegacyInterfaceStatus</p>
     * Fetch legacy interface status
     */
    getLegacyInterfaceStatusRaw(requestParameters: GetLegacyInterfaceStatusRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<LegacyInterfaceStatusDetails>>;
    /**
     * API to Fetch legacy OXI Interface Processor Status by Interface Id. <p><strong>OperationId:</strong>getLegacyInterfaceStatus</p>
     * Fetch legacy interface status
     */
    getLegacyInterfaceStatus(requestParameters: GetLegacyInterfaceStatusRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<LegacyInterfaceStatusDetails>;
    /**
     * API to Fetch list of values for all OXI interface type External Systems. <p><strong>OperationId:</strong>getOXIListOfValues</p>
     * Fetch list of values details for OXI
     */
    getOXIListOfValuesRaw(requestParameters: GetOXIListOfValuesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<FetchOXIListOfValues>>;
    /**
     * API to Fetch list of values for all OXI interface type External Systems. <p><strong>OperationId:</strong>getOXIListOfValues</p>
     * Fetch list of values details for OXI
     */
    getOXIListOfValues(requestParameters: GetOXIListOfValuesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<FetchOXIListOfValues>;
    /**
     * API to Fetch Profile Match Rules by Interface Id and Hotel Id. <p><strong>OperationId:</strong>getProfileMatchRules</p>
     * Fetch profile match rules
     */
    getProfileMatchRulesRaw(requestParameters: GetProfileMatchRulesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<FetchProfileMatchRules>>;
    /**
     * API to Fetch Profile Match Rules by Interface Id and Hotel Id. <p><strong>OperationId:</strong>getProfileMatchRules</p>
     * Fetch profile match rules
     */
    getProfileMatchRules(requestParameters: GetProfileMatchRulesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<FetchProfileMatchRules>;
    /**
     * API to Create Automatic Transmission Schedules. <p><strong>OperationId:</strong>postAutomaticTransmissionSchedules</p>
     * Create Automatic Transmission Schedules
     */
    postAutomaticTransmissionSchedulesRaw(requestParameters: PostAutomaticTransmissionSchedulesOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Status>>;
    /**
     * API to Create Automatic Transmission Schedules. <p><strong>OperationId:</strong>postAutomaticTransmissionSchedules</p>
     * Create Automatic Transmission Schedules
     */
    postAutomaticTransmissionSchedules(requestParameters: PostAutomaticTransmissionSchedulesOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Status>;
    /**
     * API to Create  Interface Mapping Conversion Codes by Hotel Id, Interface Id and Conversion Code with OPERA value and External value. <p><strong>OperationId:</strong>postConversionCodeMappings</p>
     * Create Conversion Codes Mappings
     */
    postConversionCodeMappingsRaw(requestParameters: PostConversionCodeMappingsOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<CreatedConversionCodeMappings>>;
    /**
     * API to Create  Interface Mapping Conversion Codes by Hotel Id, Interface Id and Conversion Code with OPERA value and External value. <p><strong>OperationId:</strong>postConversionCodeMappings</p>
     * Create Conversion Codes Mappings
     */
    postConversionCodeMappings(requestParameters: PostConversionCodeMappingsOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<CreatedConversionCodeMappings>;
    /**
     * API to Create Interface Setup configuration. <p><strong>OperationId:</strong>postExternalInterfaceSetups</p>
     * Create Interface Setup details
     */
    postExternalInterfaceSetupsRaw(requestParameters: PostExternalInterfaceSetupsOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Status>>;
    /**
     * API to Create Interface Setup configuration. <p><strong>OperationId:</strong>postExternalInterfaceSetups</p>
     * Create Interface Setup details
     */
    postExternalInterfaceSetups(requestParameters: PostExternalInterfaceSetupsOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Status>;
    /**
     * API to Create Profile Match Rules. <p><strong>OperationId:</strong>postProfileMatchRules</p>
     * Create profile match rules
     */
    postProfileMatchRulesRaw(requestParameters: PostProfileMatchRulesOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Status>>;
    /**
     * API to Create Profile Match Rules. <p><strong>OperationId:</strong>postProfileMatchRules</p>
     * Create profile match rules
     */
    postProfileMatchRules(requestParameters: PostProfileMatchRulesOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Status>;
    /**
     * API to Modify  Automatic Transmission Schedules by Hotel Id. <p><strong>OperationId:</strong>putAutomaticTransmissionSchedules</p>
     * Change Automatic Transmission Schedules
     */
    putAutomaticTransmissionSchedulesRaw(requestParameters: PutAutomaticTransmissionSchedulesOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChangedAutomaticTransmissionSchedules>>;
    /**
     * API to Modify  Automatic Transmission Schedules by Hotel Id. <p><strong>OperationId:</strong>putAutomaticTransmissionSchedules</p>
     * Change Automatic Transmission Schedules
     */
    putAutomaticTransmissionSchedules(requestParameters: PutAutomaticTransmissionSchedulesOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChangedAutomaticTransmissionSchedules>;
    /**
     * API to Modify Communication Methods. <p><strong>OperationId:</strong>putCommunicationMethods</p>
     * Modify Communication Methods
     */
    putCommunicationMethodsRaw(requestParameters: PutCommunicationMethodsOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChangedCommunicationMethods>>;
    /**
     * API to Modify Communication Methods. <p><strong>OperationId:</strong>putCommunicationMethods</p>
     * Modify Communication Methods
     */
    putCommunicationMethods(requestParameters: PutCommunicationMethodsOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChangedCommunicationMethods>;
    /**
     * API to Modify  Interface Mapping  Conversion Codes by Hotel Id, Interface Id and Conversion Code Id with OPERA value and External value. <p><strong>OperationId:</strong>putConversionCodeMappings</p>
     * Change Conversion Codes Mappings
     */
    putConversionCodeMappingsRaw(requestParameters: PutConversionCodeMappingsOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChangedConversionCodeMappings>>;
    /**
     * API to Modify  Interface Mapping  Conversion Codes by Hotel Id, Interface Id and Conversion Code Id with OPERA value and External value. <p><strong>OperationId:</strong>putConversionCodeMappings</p>
     * Change Conversion Codes Mappings
     */
    putConversionCodeMappings(requestParameters: PutConversionCodeMappingsOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChangedConversionCodeMappings>;
    /**
     * API to Modify Interface Mapping UDF Conversion Codes. <p><strong>OperationId:</strong>putConversionCodesStatus</p>
     * Change Conversion codes Status
     */
    putConversionCodesStatusRaw(requestParameters: PutConversionCodesStatusOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChangedConversionCodesStatus>>;
    /**
     * API to Modify Interface Mapping UDF Conversion Codes. <p><strong>OperationId:</strong>putConversionCodesStatus</p>
     * Change Conversion codes Status
     */
    putConversionCodesStatus(requestParameters: PutConversionCodesStatusOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChangedConversionCodesStatus>;
    /**
     * API to Modify  Interface Setup configuration. <p><strong>OperationId:</strong>putExternalInterfaceSetups</p>
     * Modify external Interface Setup details
     */
    putExternalInterfaceSetupsRaw(requestParameters: PutExternalInterfaceSetupsOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChangedExternalInterfaceSetups>>;
    /**
     * API to Modify  Interface Setup configuration. <p><strong>OperationId:</strong>putExternalInterfaceSetups</p>
     * Modify external Interface Setup details
     */
    putExternalInterfaceSetups(requestParameters: PutExternalInterfaceSetupsOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChangedExternalInterfaceSetups>;
    /**
     * API to Modify Interface Controls by Interface Id. <p><strong>OperationId:</strong>putInterfaceControls</p>
     * Change OXI Parameters and Defaults
     */
    putInterfaceControlsRaw(requestParameters: PutInterfaceControlsOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChangedInterfaceControls>>;
    /**
     * API to Modify Interface Controls by Interface Id. <p><strong>OperationId:</strong>putInterfaceControls</p>
     * Change OXI Parameters and Defaults
     */
    putInterfaceControls(requestParameters: PutInterfaceControlsOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChangedInterfaceControls>;
    /**
     * API to Modify Profile Match Rules. <p><strong>OperationId:</strong>putProfileMatchRules</p>
     * Change profile match rules
     */
    putProfileMatchRulesRaw(requestParameters: PutProfileMatchRulesOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChangedProfileMatchRules>>;
    /**
     * API to Modify Profile Match Rules. <p><strong>OperationId:</strong>putProfileMatchRules</p>
     * Change profile match rules
     */
    putProfileMatchRules(requestParameters: PutProfileMatchRulesOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChangedProfileMatchRules>;
}
/**
 * @export
 */
export declare const GetAccumulatedBusinessEventsModulesEnum: {
    readonly Availability: "Availability";
    readonly Block: "Block";
    readonly BlockDetail: "BlockDetail";
    readonly BlockHeader: "BlockHeader";
    readonly Rate: "Rate";
    readonly Reservation: "Reservation";
    readonly Restriction: "Restriction";
};
export type GetAccumulatedBusinessEventsModulesEnum = typeof GetAccumulatedBusinessEventsModulesEnum[keyof typeof GetAccumulatedBusinessEventsModulesEnum];
/**
 * @export
 */
export declare const GetAccumulatedBusinessEventsStatusEnum: {
    readonly New: "New";
    readonly Ready: "Ready";
};
export type GetAccumulatedBusinessEventsStatusEnum = typeof GetAccumulatedBusinessEventsStatusEnum[keyof typeof GetAccumulatedBusinessEventsStatusEnum];
/**
 * @export
 */
export declare const GetAccumulatedBusinessEventsIntegrationSystemEnum: {
    readonly Central: "Central";
    readonly Property: "Property";
    readonly Both: "Both";
};
export type GetAccumulatedBusinessEventsIntegrationSystemEnum = typeof GetAccumulatedBusinessEventsIntegrationSystemEnum[keyof typeof GetAccumulatedBusinessEventsIntegrationSystemEnum];
/**
 * @export
 */
export declare const GetConversionCodeMappingsIntegrationSystemEnum: {
    readonly Central: "Central";
    readonly Property: "Property";
    readonly Both: "Both";
};
export type GetConversionCodeMappingsIntegrationSystemEnum = typeof GetConversionCodeMappingsIntegrationSystemEnum[keyof typeof GetConversionCodeMappingsIntegrationSystemEnum];
/**
 * @export
 */
export declare const GetConversionCodeMappingsSearchByEnum: {
    readonly OperaValue: "OperaValue";
    readonly ExternalValue: "ExternalValue";
    readonly Both: "Both";
};
export type GetConversionCodeMappingsSearchByEnum = typeof GetConversionCodeMappingsSearchByEnum[keyof typeof GetConversionCodeMappingsSearchByEnum];
/**
 * @export
 */
export declare const GetConversionCodesIntegrationSystemEnum: {
    readonly Central: "Central";
    readonly Property: "Property";
    readonly Both: "Both";
};
export type GetConversionCodesIntegrationSystemEnum = typeof GetConversionCodesIntegrationSystemEnum[keyof typeof GetConversionCodesIntegrationSystemEnum];
/**
 * @export
 */
export declare const GetExternalInterfaceSetupsSystemTypeEnum: {
    readonly Central: "Central";
    readonly Property: "Property";
    readonly Both: "Both";
};
export type GetExternalInterfaceSetupsSystemTypeEnum = typeof GetExternalInterfaceSetupsSystemTypeEnum[keyof typeof GetExternalInterfaceSetupsSystemTypeEnum];
/**
 * @export
 */
export declare const GetIntegrationInboundMessagesErrorMessageTypeEnum: {
    readonly Error: "Error";
    readonly ResultException: "ResultException";
    readonly OptionalWarning: "OptionalWarning";
    readonly Warning: "Warning";
};
export type GetIntegrationInboundMessagesErrorMessageTypeEnum = typeof GetIntegrationInboundMessagesErrorMessageTypeEnum[keyof typeof GetIntegrationInboundMessagesErrorMessageTypeEnum];
/**
 * @export
 */
export declare const GetIntegrationOutboundMessagesErrorMessageTypeEnum: {
    readonly Error: "Error";
    readonly ResultException: "ResultException";
    readonly OptionalWarning: "OptionalWarning";
    readonly Warning: "Warning";
};
export type GetIntegrationOutboundMessagesErrorMessageTypeEnum = typeof GetIntegrationOutboundMessagesErrorMessageTypeEnum[keyof typeof GetIntegrationOutboundMessagesErrorMessageTypeEnum];
/**
 * @export
 */
export declare const GetInterfaceControlsIntegrationSystemEnum: {
    readonly Central: "Central";
    readonly Property: "Property";
    readonly Both: "Both";
};
export type GetInterfaceControlsIntegrationSystemEnum = typeof GetInterfaceControlsIntegrationSystemEnum[keyof typeof GetInterfaceControlsIntegrationSystemEnum];
/**
 * @export
 */
export declare const GetOXIListOfValuesIntegrationSystemEnum: {
    readonly Central: "Central";
    readonly Property: "Property";
    readonly Both: "Both";
};
export type GetOXIListOfValuesIntegrationSystemEnum = typeof GetOXIListOfValuesIntegrationSystemEnum[keyof typeof GetOXIListOfValuesIntegrationSystemEnum];
/**
 * @export
 */
export declare const GetOXIListOfValuesIncludeInActiveEnum: {
    readonly True: true;
    readonly False: false;
};
export type GetOXIListOfValuesIncludeInActiveEnum = typeof GetOXIListOfValuesIncludeInActiveEnum[keyof typeof GetOXIListOfValuesIncludeInActiveEnum];
