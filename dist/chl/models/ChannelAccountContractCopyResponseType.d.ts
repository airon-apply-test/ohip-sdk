/**
 * OPERA Cloud Channel Configuration API
 * APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ChannelAccountContractCopyResponseTypeContractResponseInformationInner } from './ChannelAccountContractCopyResponseTypeContractResponseInformationInner';
import type { ErrorType } from './ErrorType';
import type { UniqueIDType } from './UniqueIDType';
/**
 * Channel account contract type to hold account and contract details to copy.
 * @export
 * @interface ChannelAccountContractCopyResponseType
 */
export interface ChannelAccountContractCopyResponseType {
    /**
     *
     * @type {UniqueIDType}
     * @memberof ChannelAccountContractCopyResponseType
     */
    accountId?: UniqueIDType;
    /**
     * Target Account code to which contract is copied.
     * @type {string}
     * @memberof ChannelAccountContractCopyResponseType
     */
    accountCode?: string;
    /**
     * Holds contract Id of the contract copied.
     * @type {Array<ChannelAccountContractCopyResponseTypeContractResponseInformationInner>}
     * @memberof ChannelAccountContractCopyResponseType
     */
    contractResponseInformation?: Array<ChannelAccountContractCopyResponseTypeContractResponseInformationInner>;
    /**
     *
     * @type {ErrorType}
     * @memberof ChannelAccountContractCopyResponseType
     */
    errorMessage?: ErrorType;
}
/**
 * Check if a given object implements the ChannelAccountContractCopyResponseType interface.
 */
export declare function instanceOfChannelAccountContractCopyResponseType(value: object): boolean;
export declare function ChannelAccountContractCopyResponseTypeFromJSON(json: any): ChannelAccountContractCopyResponseType;
export declare function ChannelAccountContractCopyResponseTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChannelAccountContractCopyResponseType;
export declare function ChannelAccountContractCopyResponseTypeToJSON(value?: ChannelAccountContractCopyResponseType | null): any;
