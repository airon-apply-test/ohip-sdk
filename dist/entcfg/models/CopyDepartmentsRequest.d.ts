/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { CopyConfigurationCodeType } from './CopyConfigurationCodeType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface CopyDepartmentsRequest
 */
export interface CopyDepartmentsRequest {
    /**
     * List of the departments to be copied.
     * @type {Array<CopyConfigurationCodeType>}
     * @memberof CopyDepartmentsRequest
     */
    departments?: Array<CopyConfigurationCodeType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof CopyDepartmentsRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof CopyDepartmentsRequest
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the CopyDepartmentsRequest interface.
 */
export declare function instanceOfCopyDepartmentsRequest(value: object): boolean;
export declare function CopyDepartmentsRequestFromJSON(json: any): CopyDepartmentsRequest;
export declare function CopyDepartmentsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): CopyDepartmentsRequest;
export declare function CopyDepartmentsRequestToJSON(value?: CopyDepartmentsRequest | null): any;
