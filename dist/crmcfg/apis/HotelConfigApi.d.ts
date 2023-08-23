/**
 * OPERA Cloud CRM Configuration API
 * APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import * as runtime from '../runtime';
import type { ChangeCommunicationTypeRequest, CommunicationTypesDetails, PostCommunicationTypeRequest, Status } from '../models';
export interface ChangeCommunicationTypeOperationRequest {
    communicationTypeCode?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    communicationTypeToBeChanged?: ChangeCommunicationTypeRequest;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface GetCommunicationTypesRequest {
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    codes?: Array<string>;
    roles?: Set<GetCommunicationTypesRolesEnum>;
    description?: string;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface PostCommunicationTypeOperationRequest {
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    communicationTypeCriteria?: PostCommunicationTypeRequest;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
export interface RemoveCommunicationTypeRequest {
    communicationTypeCode?: string;
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    xExternalsystem?: string;
    acceptLanguage?: string;
}
/**
 *
 */
export declare class HotelConfigApi extends runtime.BaseAPI {
    /**
     * Use this API to update a  communication type. <p><strong>OperationId:</strong>changeCommunicationType</p>
     * Change a  communication type
     */
    changeCommunicationTypeRaw(requestParameters: ChangeCommunicationTypeOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Status>>;
    /**
     * Use this API to update a  communication type. <p><strong>OperationId:</strong>changeCommunicationType</p>
     * Change a  communication type
     */
    changeCommunicationType(requestParameters: ChangeCommunicationTypeOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Status>;
    /**
     * Use this API to get communication types. <p><strong>OperationId:</strong>getCommunicationTypes</p>
     * Get communication types
     */
    getCommunicationTypesRaw(requestParameters: GetCommunicationTypesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<CommunicationTypesDetails>>;
    /**
     * Use this API to get communication types. <p><strong>OperationId:</strong>getCommunicationTypes</p>
     * Get communication types
     */
    getCommunicationTypes(requestParameters: GetCommunicationTypesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<CommunicationTypesDetails>;
    /**
     * Use this API to create a  communication type. <p><strong>OperationId:</strong>postCommunicationType</p>
     * Create a  communication type
     */
    postCommunicationTypeRaw(requestParameters: PostCommunicationTypeOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Status>>;
    /**
     * Use this API to create a  communication type. <p><strong>OperationId:</strong>postCommunicationType</p>
     * Create a  communication type
     */
    postCommunicationType(requestParameters: PostCommunicationTypeOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Status>;
    /**
     * Use this API to delete a communication type. <p><strong>OperationId:</strong>removeCommunicationType</p>
     * Delete a communication type
     */
    removeCommunicationTypeRaw(requestParameters: RemoveCommunicationTypeRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Status>>;
    /**
     * Use this API to delete a communication type. <p><strong>OperationId:</strong>removeCommunicationType</p>
     * Delete a communication type
     */
    removeCommunicationType(requestParameters: RemoveCommunicationTypeRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Status>;
}
/**
 * @export
 */
export declare const GetCommunicationTypesRolesEnum: {
    readonly Email: "Email";
    readonly Fax: "Fax";
    readonly EmailForRequest: "EmailForRequest";
    readonly Webpage: "Webpage";
    readonly Phone: "Phone";
};
export type GetCommunicationTypesRolesEnum = typeof GetCommunicationTypesRolesEnum[keyof typeof GetCommunicationTypesRolesEnum];
